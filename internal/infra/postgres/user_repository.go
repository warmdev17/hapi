package postgres

import (
	"context"

	"github.com/warmdev17/hapi/internal/repositories"
	"github.com/warmdev17/hapi/internal/user"
)

type userRepoImpl struct {
	q repositories.Querier
}

func NewUserUserRepository(q repositories.Querier) user.Repository {
	return &userRepoImpl{
		q: q,
	}
}

func (r *userRepoImpl) CreateUser(ctx context.Context, args user.CreateUser) (user.User, error) {
	sqlcArgs := repositories.CreateUserParams{
		Username:       args.Username,
		Email:          args.Email,
		HashedPassword: args.HashedPassword,
		DisplayName:    args.DisplayName,
		BirthDay:       args.BirthDay,
		Gender:         string(args.Gender),
	}

	u, err := r.q.CreateUser(ctx, sqlcArgs)
	return user.User{
		ID:          u.ID,
		Username:    u.Username,
		Email:       u.Email,
		DisplayName: u.DisplayName,
		BirthDay:    u.BirthDay,
		Gender:      u.Gender,
	}, err

}

func (r *userRepoImpl) GetUserByEmail(ctx context.Context, email string) (user.User, error) {
	u, err := r.q.GetUserByEmail(ctx, email)
	return user.User{
		ID:          u.ID,
		Username:    u.Username,
		DisplayName: u.DisplayName,
		Email:       u.Email,
		BirthDay:    u.BirthDay,
		Gender:      u.Gender,
	}, err
}

func (r *userRepoImpl) CheckEmailExists(ctx context.Context, email string) (bool, error) {
	return r.q.CheckEmailExists(ctx, email)
}
