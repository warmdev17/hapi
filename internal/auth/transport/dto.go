package transport

import (
	"time"

	"github.com/google/uuid"
)

type RegisterRequest struct {
	Username    string    `json:"username" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required,min=8"`
	DisplayName *string   `json:"displayName"`
	BirthDay    time.Time `json:"birthDay" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
}

type RegisterRespons struct {
	AccessToken string       `json:"accessToken"`
	User        UserResponse `json:"user"`
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
}
