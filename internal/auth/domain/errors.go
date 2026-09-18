package domain

import "errors"

var (
	ErrEmailTaken    error = errors.New("email already taken")
	ErrUsernameTaken error = errors.New("username already taken")
)
