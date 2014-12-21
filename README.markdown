# tagit

[![BuildStatus](https://travis-ci.org/ndyakov/tagit.png)](https://travis-ci.org/ndyakov/tagit)
[![GoDoc](https://godoc.org/github.com/ndyakov/tagit?status.png)](https://godoc.org/github.com/ndyakov/tagit)

__tagit__ is a package that exports two types
that can be used to add tag lists to your structs.

### Taggable

The first of them (`tagit.Taggable`) can be used as anonymous field
in a struct:

```go
	type Article {
		tagit.Taggable
	}
```
This type provides barebones tag operations and cannot be Marshaled/Unmarshaled by itself.

### Tags

The other one (`tagit.Tags`) is designed to be used as an field in
a struct (as composite) and is the prefered way of using taggit:

```go
	type Article {
		Tags *tagit.Tags `json:"tags"`
	}
```

By using tagit.Tags you will be able to use `json.Marshal` and `json.Unmarshal`
on your type.

When using `tagit.Tags` you will have to initialize it with the `tagit.NewTags()` function.

## Usage

```
go get github.com/ndyakov/tagit
```

```go
import "github.com/ndyakov/tagit"
```

## Tagit + mgo

```go
import "github.com/ndyakov/tagit/bson"
```

If you want to use this with mgo (the mongo driver for golang) you can use the tagit/bson packet
that has a `Tags` type similar to the one in the root packet but with the possibility to be
Marshaled to BSON and Unmarshaled from BSON. This can be used with mgo.

```go
	type Artwork {
		Name string      `bson:"name, omitempty" json:"name"`
		Tags *tagit.Tags `bson:"tags, omitempty" json:"tags"`
	}

```

Check the [Godoc](https://godoc.org/github.com/ndyakov/tagit) and decide which type you want to use.
