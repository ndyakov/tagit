/*
Tagit is a package that exports two types
that can be used to add tag lists to your structs.
The first of them (tagit.Taggable) can be used as anonymous field
in a struct:

	type Article {
		tagit.Taggable
	}

The other one (tagit.Tags) is designed to be used as an field in
a struct (composition) and is the prefered way of using taggit:

	type Article {
		Tags *tagit.Tags `json:"tags"`
	}

By using tagit.Tags you will be able to use json.Marshal and json.Unmarshal
on your type.

When using tagit.Tags you will have to initialize it with the tagit.NewTags() function.
*/

package tagit
