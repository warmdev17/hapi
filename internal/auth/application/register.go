package application

import (
	"context"
	"time"

	"github.com/google/uuid"
	authdomain "github.com/warmdev17/hapi/internal/auth/domain"
	userdomain "github.com/warmdev17/hapi/internal/user/domain"
)

type RegisterInput struct {
	Username    string
	Email       string
	Password    string
	DisplayName string
	BirthDay    time.Time
	Gender      string
}

type UserCreateResult struct {
	ID          uuid.UUID
	Username    string
	DisplayName string
}

type RegisterResult struct {
	AccessToken string
	User        UserCreateResult
}

type UserRepository interface {
	ExistsByEmail(ctx context.Context, email userdomain.Email) (bool, error)
	ExistsByUsername(ctx context.Context, username string) (bool, error)
	Save(ctx context.Context, user *userdomain.User) (*userdomain.User, error)
}

type PasswordHasher interface {
	Hash(password string) (string, error)
}

type TokenProvider interface {
	Generate(userID uuid.UUID, ttl time.Duration) (string, error)
}

type RegisterService struct {
	userRepo       UserRepository
	passwordHasher PasswordHasher
	tokenPvd       TokenProvider
	accessTTL      time.Duration
	// TODO: refresh token
	// refreshTTL     time.Duration
}

func NewRegisterService(
	userRepo UserRepository,
	passwordHasher PasswordHasher,
	tokenPvd TokenProvider,
	accessTTL time.Duration,
	// refreshTTL time.Duration,
) *RegisterService {
	return &RegisterService{
		userRepo:       userRepo,
		passwordHasher: passwordHasher,
		tokenPvd:       tokenPvd,
		accessTTL:      accessTTL,
		// refreshTTL:     refreshTTL,
	}
}

func (s *RegisterService) Execute(
	ctx context.Context,
	input RegisterInput,
) (*RegisterResult, error) {
	email, err := userdomain.NewEmail(input.Email)
	if err != nil {
		return nil, err
	}

	emailTaken, err := s.userRepo.ExistsByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if emailTaken {
		return nil, authdomain.ErrEmailTaken
	}

	usernameTaken, err := s.userRepo.ExistsByUsername(ctx, input.Username)
	if err != nil {
		return nil, err
	}
	if usernameTaken {
		return nil, authdomain.ErrUsernameTaken
	}

	gender, err := userdomain.NewGender(input.Gender)
	if err != nil {
		return nil, err
	}

	hashedPassword, err := s.passwordHasher.Hash(input.Password)
	if err != nil {
		return nil, err
	}

	user, err := userdomain.NewUser(
		input.Username,
		email,
		hashedPassword,
		input.DisplayName,
		input.BirthDay,
		gender,
	)
	if err != nil {
		return nil, err
	}

	createdUser, err := s.userRepo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	accessToken, err := s.tokenPvd.Generate(createdUser.ID(), s.accessTTL)
	if err != nil {
		return nil, err
	}

	return &RegisterResult{
		AccessToken: accessToken,
		User: UserCreateResult{
			ID:          createdUser.ID(),
			Username:    createdUser.Username(),
			DisplayName: createdUser.DisplayName(),
		},
	}, nil
}
