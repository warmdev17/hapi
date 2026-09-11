package auth

import (
	"time"

	"github.com/warmdev17/hapi/internal/user"
)

type Gender string

const (
	Male   Gender = "male"
	Female Gender = "female"
	Other  Gender = "other"
)

type RegisterInput struct {
	Username        string    `json:"username" binding:"required,min=1,max=255"`
	Email           string    `json:"email" binding:"required,email,min=1,max=255"`
	Password        string    `json:"password" binding:"required,min=8,max=32"`
	DisplayName     string    `json:"displayName" binding:"required,min=1,max=255"`
	ConfirmPassword string    `json:"confirmPassword" binding:"required"`
	BirthDay        time.Time `json:"birthDay" binding:"required"`
	Gender          Gender    `json:"gender" binding:"required,oneof=male female other"`
}

type LoginInput struct {
	Identifier string `json:"identifier" binding:"required"`
	Password   string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string            `json:"accessToken"`
	User        user.UserResponse `json:"user"`
}
