# tagit

[![BuildStatus](https://travis-ci.org/ndyakov/tagit.png)](https://travis-ci.org/ndyakov/tagit)
[![GoDoc](https://godoc.org/github.com/ndyakov/tagit?status.png)](https://godoc.org/github.com/ndyakov/tagit)

__tagit__ is a package that exports two types
that can be used to add tag lists to your structs.

### Taggable

__This may be removed, because I don't see why anyone whould like to use it instead of `Tags`__

The first of them (`tagit.Taggable`) can be used as anonymous field
in a struct:

```go
	type Article struct {
		tagit.Taggable
	}
```
This type provides barebones tag operations and cannot be Marshaled/Unmarshaled by itself.


### Tags

The other one (`tagit.Tags`) is designed to be used as an field in
a struct (as composite) and is the prefered way of using taggit:

```go
	type Article struct {
		Tags *tagit.Tags `json:"tags"`
	}
```

By using tagit.Tags you will be able to use `json.Marshal` and `json.Unmarshal`
on your type.

When using `tagit.Tags` you will have to initialize it with the `tagit.NewTags()` function.

```
type Tags
  func NewTags() *Tags
  func (t *Tags) Add(tag string)
  func (t *Tags) All() []string
  func (t *Tags) Count() int
  func (t *Tags) Has(tag string) bool
  func (t *Tags) MarshalJSON() ([]byte, error)
  func (t *Tags) Remove(tag string)
  func (t *Tags) String() (res string)
  func (t *Tags) UnmarshalJSON(json []byte) error
```

## Usage

```
go get github.com/ndyakov/tagit
```

```go

type Article struct {
	Name string      `json:"name"`
	Tags *tagit.Tags `json:"tags"`
}

func NewArticle(name) *Article {
	return &Article{Name: name, Tags: tagit.NewTags()}
}

func main() {
	a := NewArticle("Tagit!")
	a.Tags.Add("example")
	tags := a.Tags.All()  // ['example']
	a.Tags.Has("example") // true
	a.Tags.Count()        // 1
}

```

## Tagit + bson
[![GoDoc](https://godoc.org/github.com/ndyakov/tagit/bson?status.png)](https://godoc.org/github.com/ndyakov/tagit/bson)

```go
import "github.com/ndyakov/tagit/bson"
```

If you want to use this with mgo (the mongo driver for golang) you can use the tagit/bson packet
that has a `Tags` type similar to the one in the root packet but with the possibility to be
Marshaled to BSON and Unmarshaled from BSON. This can be used with mgo.

```go
	type Artwork struct {
		Name string      `bson:"name, omitempty" json:"name"`
		Tags *tagit.Tags `bson:"tags, omitempty" json:"tags"`
	}

```


#### Why are there different packages with almost the same type?

You may wonder why there are few different packages with almost the same `Tags` type
an which one you should use.

The main reason is that working with bson, for example, needs additional imports and I want to
keep the root package as slim as possible. Althought this will be harder to maintain I think it is a
reasonable solutions for small package as this one.
