package tagit

// Taggable is a simple type that you can add as
// anonymous field to your struct and then tag/untag
// instances of your type
type Taggable struct {
	tags map[string]struct{}
}

func (t *Taggable) initTaggable() {
	if t.tags == nil {
		t.tags = make(map[string]struct{})
	}
}

// Tags returns the tags list as slice of strings.
// Be aware that the order of the elements may differ between calls.
func (t *Taggable) Tags() []string {
	t.initTaggable()
	return keys(t.tags)
}

// HasTag checks if a tag is available in the tag list.
func (t *Taggable) HasTag(tag string) bool {
	t.initTaggable()
	_, ok := t.tags[tag]
	return ok
}

// Tag adds tag to the list.
func (t *Taggable) Tag(tag string) {
	t.initTaggable()
	t.tags[tag] = struct{}{}
}

// Untag removes tag from the list.
func (t *Taggable) Untag(tag string) {
	t.initTaggable()
	delete(t.tags, tag)
}

// NumTags counts the number of tags in the list.
func (t *Taggable) NumTags() int {
	return len(t.Tags())
}
