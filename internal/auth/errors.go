package auth

import "errors"

var (
	ErrEmailTaken = errors.New("email already registered")
	ErrInternal   = errors.New("an unexpected error occured while processing your request")
)
