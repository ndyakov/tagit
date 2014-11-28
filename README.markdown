# tagit

[![BuildStatus](https://travis-ci.org/ndyakov/tagit.png)](https://travis-ci.org/ndyakov/tagit)
[![GoDoc](https://godoc.org/github.com/ndyakov/tagit?status.png)](https://godoc.org/github.com/ndyakov/tagit)

__tagit__ is a package that exports two types
that can be used to add tag lists to your structs.
The first of them (`tagit.Taggable`) can be used as anonymous field
in a struct:

```go
	type Article {
		tagit.Taggable
	}
```

The other one (`tagit.Tags`) is designed to be used as an field in
a struct (composition) and is the prefered way of using taggit:

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

Check the [Godoc](https://godoc.org/github.com/ndyakov/tagit) and decide wich type you want to use.
