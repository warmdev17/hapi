package transport

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	authapplication "github.com/warmdev17/hapi/internal/auth/application"
	"github.com/warmdev17/hapi/pkg/response"
)

type RegisterApplication interface {
	Execute(ctx context.Context, input authapplication.RegisterInput) (*authapplication.RegisterResult, error)
}

type LoginApplication interface {
	Execute(ctx context.Context, input LoginRequest)
}

type Handler struct {
	register RegisterApplication
	// login    LoginApplication
}

func NewHandler(
	register RegisterApplication,
	// login LoginApplication,
) *Handler {
	return &Handler{
		register: register,
		// login:    login,
	}
}

func (h *Handler) Register(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		status, code, message := MapError(err)
		response.Error(c, status, code, message)
	}

	input := authapplication.RegisterInput{
		Username:    req.Username,
		Email:       req.Email,
		Password:    req.Password,
		DisplayName: *req.DisplayName,
		BirthDay:    req.BirthDay,
		Gender:      req.Gender,
	}
	result, err := h.register.Execute(c.Request.Context(), input)
	if err != nil {
		status, code, message := MapError(err)
		response.Error(c, status, code, message)
	}

	response.Success(c, http.StatusCreated, ToRegisterResponse(*result))

}

func (h *Handler) Login(c *gin.Context) {
	fmt.Print("hello world")
}
