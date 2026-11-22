package model

import "errors"

var ErrFooNotFound = errors.New("foo not found")

type Foo struct {
	ID  int64  `json:"id"`
	Foo string `json:"foo"`
}
