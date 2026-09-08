package auth

import "errors"

var (
	ErrEmailTaken = errors.New("email already registered")
)
