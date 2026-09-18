package transport

import (
	"time"

	"github.com/google/uuid"
	"github.com/warmdev17/hapi/internal/auth/application"
)

type RegisterRequest struct {
	Username    string    `json:"username" binding:"required"`
	Email       string    `json:"email" binding:"required"`
	Password    string    `json:"password" binding:"required,min=8"`
	DisplayName *string   `json:"displayName"`
	BirthDay    time.Time `json:"birthDay" binding:"required"`
	Gender      string    `json:"gender" binding:"required"`
}

type RegisterResponse struct {
	AccessToken string       `json:"accessToken"`
	User        UserResponse `json:"user"`
}

type LoginRequest struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID          uuid.UUID `json:"id"`
	Username    string    `json:"username"`
	DisplayName string    `json:"displayName"`
}

func ToRegisterResponse(res application.RegisterResult) RegisterResponse {
	return RegisterResponse{
		AccessToken: res.AccessToken,
		User: UserResponse{
			ID:          res.User.ID,
			Username:    res.User.Username,
			DisplayName: res.User.DisplayName,
		},
	}
}
