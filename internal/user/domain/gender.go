package domain

import (
	"strings"
)

type Gender string

const (
	GenderMale   Gender = "male"
	GenderFemale Gender = "female"
	GenderOther  Gender = "other"
)

func NewGender(value string) (Gender, error) {
	value = strings.TrimSpace(strings.ToLower(value))

	switch Gender(value) {
	case GenderMale, GenderFemale, GenderOther:
		return Gender(value), nil
	default:
		return "", ErrInvalidGender
	}
}
