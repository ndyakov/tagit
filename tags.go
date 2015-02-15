package tagit

import (
	"errors"
	"strconv"
	"strings"
)

// Tags struct that holds tags and have a simple interface for working with
// the tags. By composing this the list of tags will be able
// to be marshalled/unmarshalled to/from JSON.
type Tags struct {
	tags map[string]struct{}
}

// NewTags is a constructor for Tags.
// You should initialize your tags with this function.
// This function can accept variable number of string agruments
// as an initial set of tags for the created object.
func NewTags(tags ...string) *Tags {
	t := new(Tags)
	t.initTags()

	for _, tag := range tags {
		t.Add(tag)
	}

	return t
}

func (t *Tags) initTags() {
	if t.tags == nil {
		t.tags = make(map[string]struct{})
	}
}

// All returns all tags as a slice of strings.
// The order of the tags in the slice may vary between different calls.
func (t *Tags) All() []string {
	t.initTags()
	return keys(t.tags)
}

// Has checks if a tags is in the list.
// Returns boolean.
func (t *Tags) Has(tag string) bool {
	t.initTags()
	_, ok := t.tags[strings.TrimSpace(tag)]
	return ok
}

// Add adds a tag to the list.
func (t *Tags) Add(tag string) {
	t.initTags()
	t.tags[strings.TrimSpace(tag)] = struct{}{}
}

// Remove removes a tag from the list.
func (t *Tags) Remove(tag string) {
	t.initTags()
	delete(t.tags, strings.TrimSpace(tag))
}

// Count counts the number of tags and return it as int.
func (t *Tags) Count() int {
	return len(t.All())
}

// String to implement fmt.Stringer.
// Return tags as comma separated list of words.
func (t *Tags) String() (res string) {
	if len(t.All()) <= 0 {
		return
	}

	for _, tag := range t.All() {
		res += ", " + tag
	}

	return res[2:]
}

func (t *Tags) quotedTags() (res string) {
	if len(t.All()) <= 0 {
		return
	}

	for _, tag := range t.All() {
		res += "," + strconv.Quote(tag)
	}

	return res[1:]
}

// MarshalJSON to implement json.Marshaller.
// Returns the tags as JSON list of strings.
func (t *Tags) MarshalJSON() ([]byte, error) {
	return []byte("[" + t.quotedTags() + "]"), nil
}

// UnmarshalJSON  to implement json.Unmarshaller.
// Expects JSON list of strings. Adds the tags to the list.
// Will not remove the current tags in the list.
func (t *Tags) UnmarshalJSON(json []byte) error {
	var quote, backSlash, leftBracket, rightBracket byte
	quote, leftBracket, backSlash, rightBracket = 34, 91, 92, 93
	if json[0] != leftBracket || json[len(json)-1] != rightBracket {
		return errors.New("provided JSON is not a list")
	}
	var wordBytes []byte
	inWord := false
	for pos := 1; pos < len(json); pos++ {
		if inWord {
			if json[pos] == backSlash {
				slashesh := 1
				pos++
				for json[pos] == backSlash {
					slashesh++
					if slashesh == 2 {
						wordBytes = append(wordBytes, backSlash)
						slashesh = 0
					}
					pos++
				}
				if slashesh == 0 {
					pos--
					continue
				}
			} else if json[pos] == quote {
				inWord = false
				t.Add(string(wordBytes))
				continue
			}
			wordBytes = append(wordBytes, json[pos])
		} else {
			if json[pos] == quote {
				wordBytes = make([]byte, 0, len(json)-pos)
				inWord = true
			}
		}
	}
	return nil
}
