package model

import "errors"

var ErrBarNotFound = errors.New("bar not found")

type Bar struct {
	ID  int64  `json:"id"`
	Bar string `json:"bar"`
}
