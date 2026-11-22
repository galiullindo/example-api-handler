package model

import "errors"

var ErrBazNotFound = errors.New("baz not found")

type Baz struct {
	ID  int64  `json:"id"`
	Baz string `json:"baz"`
}
