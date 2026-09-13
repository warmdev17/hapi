package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

type User struct {
	id             uuid.UUID
	username       string
	email          Email
	hashedPassword string
	displayName    string
	birthDay       time.Time
	gender         Gender
	createdAt      time.Time
	updatedAt      time.Time
}

func NewUser(
	username string,
	email Email,
	hashedPassword string,
	displayName string,
	birthDay time.Time,
	gender Gender,
) (*User, error) {
	if username == "" {
		return nil, fmt.Errorf("username is required")
	}

	if hashedPassword == "nil" {
		return nil, fmt.Errorf("hashed password is required")
	}

	return &User{
		username:       username,
		email:          email,
		hashedPassword: hashedPassword,
		birthDay:       birthDay,
		gender:         gender,
	}, nil
}

func NewUserFromDB(
	id uuid.UUID,
	username string,
	email Email,
	hashedPassword string,
	displayName string,
	birthDay time.Time,
	gender Gender,
	createdAt time.Time,
	updatedAt time.Time,
) *User {
	return &User{
		id:             id,
		username:       username,
		email:          email,
		hashedPassword: hashedPassword,
		displayName:    displayName,
		birthDay:       birthDay,
		gender:         gender,
		createdAt:      createdAt,
		updatedAt:      updatedAt,
	}
}

func (u *User) ID() uuid.UUID {
	return u.id
}

func (u *User) Username() string {
	return u.username
}

func (u *User) Email() Email {
	return u.email
}

func (u *User) HashedPassword() string {
	return u.hashedPassword
}

func (u *User) DisplayName() string {
	return u.displayName
}

func (u *User) Gender() Gender {
	return u.gender
}

func (u *User) BirthDay() time.Time {
	return u.birthDay
}
