/*
Package tagit/bson is a package that exports a type
that can be used to add tag lists to your structs.

The tagit.Tags is designed to be used as an field in
a struct (composition) like so:

	type Article {
		Tags *tagit.Tags `json:"tags"`
	}

By using tagit.Tags you will be able to use json.Marshal and json.Unmarshal
on your type.

You will be able to use mgo/bson with this package. The type tagit.Tags knows how to
be marshalled and unmarshalled to/from bson.

When using tagit.Tags you will have to initialize it with the tagit.NewTags() function.
*/
package tagit
