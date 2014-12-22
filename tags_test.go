package tagit_test

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/ndyakov/tagit"
)

type Comment struct {
	Tags *tagit.Tags `json:"tags"`
}

func NewComment() *Comment {
	c := new(Comment)
	c.Tags = tagit.NewTags()
	return c
}

func equalSlices(expected, got []string) bool {
	if len(expected) != len(got) {
		return false
	}

	for ie, e := range expected {
		for ig, g := range got {
			if g == e {
				got = append(got[:ig], got[ig+1:]...)
			}
		}

		if len(got) != len(expected)-(1+ie) {
			return false
		}
	}

	if len(got) != 0 {
		return false
	}

	return true
}

func TestTagsAdd(t *testing.T) {
	c := NewComment()
	c.Tags.Add("wow")
	tags := c.Tags.All()
	expected := []string{"wow"}
	if !reflect.DeepEqual(tags, expected) {
		t.Errorf("Wrong tags! Expected %#v, but got %#v!\n", expected, tags)
	}
}

func TestTagsCount(t *testing.T) {
	c := NewComment()
	c.Tags.Add("wow")
	c.Tags.Add("such")
	c.Tags.Add("tag")
	numTags := c.Tags.Count()
	expected := 3
	if numTags != expected {
		t.Errorf("Wrong number of tags! Expected %d, but got %d!\n", expected, numTags)
	}
}

func TestTagsHas(t *testing.T) {
	c := NewComment()
	c.Tags.Add("wow")
	if !c.Tags.Has("wow") {
		t.Error("Expected to have the tag \"wow\" but it did not!")
	}

	if c.Tags.Has("such") {
		t.Error("Expected to don't have the tag \"such\" but it did!")
	}
}

func TestTagsRemove(t *testing.T) {
	c := NewComment()
	c.Tags.Add("wow")

	if !c.Tags.Has("wow") {
		t.Error("Expected to have the tag \"wow\" but it did not!")
	}

	c.Tags.Remove("wow")

	if c.Tags.Has("wow") {
		t.Error("Expected to don't have the tag \"wow\" but it did!")
	}
}

func TestTagsString(t *testing.T) {
	c := NewComment()
	c.Tags.Add("w\"ow")
	c.Tags.Add("such")
	got := c.Tags.String()

	expected1 := "w\"ow, such"
	expected2 := "such, w\"ow"

	if got != expected1 && got != expected2 {
		t.Errorf("Expected String() to return %s or %s but got: %s!", expected1, expected2, got)
	}
}

func TestTagsMarshalJSON(t *testing.T) {
	c := NewComment()
	c.Tags.Add("wow")
	c.Tags.Add("such")
	got, err := c.Tags.MarshalJSON()
	if err != nil {
		t.Error(err)
	}

	expected1 := []byte(`["wow","such"]`)
	expected2 := []byte(`["such","wow"]`)

	if !reflect.DeepEqual(got, expected1) && !reflect.DeepEqual(got, expected2) {
		t.Errorf("Expected MarshalJSON() to return %s or %s but got: %s!", expected1, expected2, got)
	}
}

func TestTagsMarshal(t *testing.T) {
	c := NewComment()
	c.Tags.Add("wow")
	c.Tags.Add("such")
	got, err := json.Marshal(c)
	if err != nil {
		t.Error(err)
	}

	expected1 := []byte(`{"tags":["wow","such"]}`)
	expected2 := []byte(`{"tags":["such","wow"]}`)

	if !reflect.DeepEqual(got, expected1) && !reflect.DeepEqual(got, expected2) {
		t.Errorf("Expected marshaling the comment to return %s or %s but got: %s!", expected1, expected2, got)
	}

}

func TestTagsUnmarshal(t *testing.T) {
	c := NewComment()
	jsonData := []byte(`{"tags":["such", "wow"]}`)
	err := json.Unmarshal(jsonData, &c)
	if err != nil {
		t.Error(err)
	}

	got := c.Tags.All()
	expected := []string{"wow", "such"}
	if !equalSlices(expected, got) {
		t.Errorf("Tags are not the one that we expected them to be: got: %v, expected: %v", got, expected)
	}
}

func TestTagsUnmarshalEscapedQuote(t *testing.T) {
	c := NewComment()
	jsonData := []byte(`{"tags":["suc\"h", "wow"]}`)
	err := json.Unmarshal(jsonData, &c)

	if err != nil {
		t.Error(err)
	}

	got := c.Tags.All()
	expected := []string{"wow", "suc\"h"}
	if !equalSlices(expected, got) {
		t.Errorf("Tags are not the one that we expected them to be: got: %v, expected: %v", got, expected)
	}
}

func TestTagsUnmarshalEscapedBackSlash(t *testing.T) {
	c := NewComment()
	jsonData := []byte(`{"tags":["suc\\h", "wow"]}`)
	err := json.Unmarshal(jsonData, &c)

	if err != nil {
		t.Error(err)
	}

	got := c.Tags.All()
	expected := []string{"wow", "suc\\h"}
	if !equalSlices(expected, got) {
		t.Errorf("Tags are not the one that we expected them to be: got: %v, expected: %v", got, expected)
	}
}

func TestTagsUnmarshalString(t *testing.T) {
	c := NewComment()
	jsonData := []byte(`{"tags":"such"}`)
	err := json.Unmarshal(jsonData, &c)
	if err == nil {
		t.Error("Expected to return an error when unmarshalling something else than a list.")
	}
}
