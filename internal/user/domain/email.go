package domain

import (
	"fmt"
	"strings"
)

type Email string

func NewEmail(value string) (Email, error) {
	value = strings.TrimSpace(strings.ToLower(value))

	if value == "" {
		return "", fmt.Errorf("email is required")
	}

	return Email(value), nil
}

func (e Email) String() string {
	return string(e)
}
