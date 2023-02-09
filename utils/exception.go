package utils

import "errors"

var (
	ErrId       = errors.New("id cannot be 0")
	ErrUsername = errors.New("id cannot be empty")
)
