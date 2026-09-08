package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/warmdev17/hapi/pkg/jwt"
)

type TokenProvider interface {
	GenerateToken(userID uuid.UUID, ttl time.Duration) (string, error)
	ParseToken(token string) (*jwt.Claims, error)
}

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (bool, error)
	CheckEmailExists(ctx context.Context, email string) (bool, error)
}

type AuthService struct {
	tokenPvd TokenProvider
	userRepo UserRepository
}

func NewService(provider TokenProvider) *AuthService {
	return &AuthService{tokenPvd: provider}
}

func (s *AuthService) Register(ctx context.Context, req RegisterInput) (*LoginResponse, error) {
	exists, err := s.userRepo.CheckEmailExists(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if exists {
		return nil, ErrEmailTaken
	}

}
