package apperror

import "errors"

var (
	ErrInternal = errors.New("an unexpected error occured while processing your request")
)
