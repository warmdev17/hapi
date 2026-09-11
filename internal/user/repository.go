package user

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type User struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	DisplayName string    `json:"display_name"`
	BirthDay    time.Time `json:"birth_day"`
	Gender      string    `json:"gender"`
}

type CreateUser struct {
	Username       string    `json:"username"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"hashed_password"`
	DisplayName    string    `json:"display_name"`
	BirthDay       time.Time `json:"birth_day"`
	Gender         `json:"gender"`
}

type Repository interface {
	CreateUser(ctx context.Context, args CreateUser) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	CheckEmailExists(ctx context.Context, email string) (bool, error)
}
