# go-jsonint

Package **jsonint** provides tools for working with integers in JSON.

This helps to work-around that numbers in JSON are effectively "broken".

## Documention

Online documentation, which includes examples, can be found at: http://godoc.org/sourcecode.social/reiver/go-jsonint

[![GoDoc](https://godoc.org/sourcecode.social/reiver/go-jsonint?status.svg)](https://godoc.org/sourcecode.social/reiver/go-jsonint)

## Example

Here is an example:
```go
type Account struct {
	FollowersCount jsonint.Int `json:"followers_count"`
	FollowingCount jsonint.Int `json:"following_count"`
	StatusexCount  jsonint.Int `json:"statuses_count"`
}

// ...

var account Account

err := json.Unmarshal(data, &account)
```

## Import

To import package **jsonint** use `import` code like the follownig:
```
import "sourcecode.social/reiver/go-jsonint"
```

## Installation

To install package **jsonint** do the following:
```
GOPROXY=direct https://sourcecode.social/reiver/go-jsonint
```

## Author

Package **jsonint** was written by [Charles Iliya Krempeaux](http://changelog.ca)
