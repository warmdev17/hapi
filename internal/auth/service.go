package auth

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/warmdev17/hapi/internal/user"
	"github.com/warmdev17/hapi/pkg/apperror"
	"github.com/warmdev17/hapi/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

type TokenProvider interface {
	GenerateToken(userID uuid.UUID, ttl time.Duration) (string, error)
	ParseToken(token string) (*jwt.Claims, error)
}

type UserRepository interface {
	CreateUser(ctx context.Context, args user.CreateUser) (user.User, error)
	GetUserByEmail(ctx context.Context, email string) (bool, error)
	CheckEmailExists(ctx context.Context, email string) (bool, error)
}

type AuthService struct {
	tokenPvd   TokenProvider
	userRepo   UserRepository
	accessTTL  time.Duration
	refreshTTL time.Duration
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

	hashed_password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apperror.ErrInternal
	}

	uCreateParam := user.CreateUser{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: string(hashed_password),
		DisplayName:    req.DisplayName,
		BirthDay:       req.BirthDay,
		Gender:         user.Gender(req.Gender),
	}

	u, err := s.userRepo.CreateUser(ctx, uCreateParam)

	accessToken, err := s.tokenPvd.GenerateToken(u.ID, s.accessTTL)

	return &LoginResponse{
		AccessToken: accessToken,
		User: user.UserResponse{
			ID:          u.ID,
			Username:    u.Username,
			DisplayName: u.DisplayName,
		},
	}, err

}

func (s *AuthService) Login(ctx context.Context, req LoginInput) (*LoginResponse, error)
