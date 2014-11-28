package tagit_test

import (
	tagit "."
	"reflect"
	"testing"
)

type Article struct {
	tagit.Taggable `json:"tags"`
}

func TestTaggitTag(t *testing.T) {
	a := new(Article)
	a.Tag("wow")
	tags := a.Tags()
	expected := []string{"wow"}
	if !reflect.DeepEqual(tags, expected) {
		t.Errorf("Wrong tags! Expected %#v, but got %#v!\n", expected, tags)
	}
}

func TestTaggitNumTags(t *testing.T) {
	a := new(Article)
	a.Tag("wow")
	a.Tag("such")
	a.Tag("tag")
	numTags := a.NumTags()
	expected := 3
	if numTags != expected {
		t.Errorf("Wrong number of tags! Expected %d, but got %d!\n", expected, numTags)
	}
}

func TestTaggitHasTag(t *testing.T) {
	a := new(Article)
	a.Tag("wow")
	if !a.HasTag("wow") {
		t.Error("Expected to have the tag \"wow\" but it did not!")
	}

	if a.HasTag("such") {
		t.Error("Expected to don't have the tag \"such\" but it did!")
	}
}

func TestTaggitUntag(t *testing.T) {
	a := new(Article)
	a.Tag("wow")

	if !a.HasTag("wow") {
		t.Error("Expected to have the tag \"wow\" but it did not!")
	}

	a.Untag("wow")

	if a.HasTag("wow") {
		t.Error("Expected to don't have the tag \"wow\" but it did!")
	}
}
