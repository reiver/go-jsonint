# go-jsonint

Package **jsonint** provides tools for working with integers in JSON.

This helps to work-around that numbers in JSON are effectively "broken" in many implementations, and some specifications

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/github.com/reiver/go-jsonint

[![GoDoc](https://godoc.org/github.com/reiver/go-jsonint?status.svg)](https://godoc.org/github.com/reiver/go-jsonint)

## Example

Here is an example using `jsonint.Integer`, which accepts bare JSON integers (ex: `123`):
```go
type Account struct {
	FollowersCount jsonint.Integer `json:"followers_count"`
	FollowingCount jsonint.Integer `json:"following_count"`
	StatusesCount  jsonint.Integer `json:"statuses_count"`
}

// ...

var account Account

err := json.Unmarshal(data, &account)
```

`jsonint.Numeric` is similar to `jsonint.Integer` but also accepts JSON strings containing integer values (ex: both `123` and `"123"`):
```go
type Account struct {
	FollowersCount jsonint.Numeric `json:"followers_count"`
	FollowingCount jsonint.Numeric `json:"following_count"`
	StatusesCount  jsonint.Numeric `json:"statuses_count"`
}

// ...

var account Account

err := json.Unmarshal(data, &account)
```

## Import

To import package **jsonint** use `import` code like the follownig:
```
import "github.com/reiver/go-jsonint"
```

## Installation

To install package **jsonint** do the following:
```
GOPROXY=direct go get https://github.com/reiver/go-jsonint
```

## Author

Package **jsonint** was written by [Charles Iliya Krempeaux](http://reiver.link)
