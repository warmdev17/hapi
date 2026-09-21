package domain

import (
	"unicode"
	"unicode/utf8"
)

type Password string

func NewPassword(value string) (Password, error) {
	if utf8.RuneCountInString(value) < MinPasswordLength {
		return "", ErrPasswordMinLength
	}

	strong := isPasswordStrong(value)
	if !strong {
		return "", ErrPasswordNotStrong
	}

	return Password(value), nil
}

func isPasswordStrong(pwd string) bool {
	var (
		hasNumber  bool
		hasSpecial bool
	)

	for _, r := range pwd {
		switch {
		case unicode.IsNumber(r):
			hasNumber = true
		case unicode.IsPunct(r) || unicode.IsSymbol(r):
			hasSpecial = true
		}

		if hasNumber && hasSpecial {
			return true
		}
	}

	return false
}

func (p Password) String() string {
	return "********"
}
