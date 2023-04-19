package utils

import "errors"

var (
	ErrId             = errors.New("id cannot be 0")
	ErrUsername       = errors.New("id cannot be empty")
	ErrUniqueUsername = errors.New("username already taken")
	ErrValidation     = errors.New("validation error")
)
