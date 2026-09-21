package domain

import (
	"errors"
	"fmt"
)

const MinPasswordLength = 8

var (
	ErrInvalidGender     error = errors.New("invalid gender")
	ErrPasswordMinLength error = fmt.Errorf("password must be at least %d characters", MinPasswordLength)
	ErrPasswordNotStrong error = errors.New("password must have at least 1 number and 1 character")
)
