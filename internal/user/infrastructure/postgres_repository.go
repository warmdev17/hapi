package infrastructure

import (
	"context"

	db "github.com/warmdev17/hapi/internal/db/sqlc"
	userdomain "github.com/warmdev17/hapi/internal/user/domain"
)

type PostgresUserRepository struct {
	queries *db.Queries
}

func NewPostgresUserRepository(queries *db.Queries) *PostgresUserRepository {
	return &PostgresUserRepository{
		queries: queries,
	}
}

func (r *PostgresUserRepository) ExistsByEmail(
	ctx context.Context,
	email userdomain.Email,
) (bool, error) {
	return r.queries.CheckEmailExists(ctx, email.String())
}

func (r *PostgresUserRepository) ExistsByUsername(
	ctx context.Context,
	username string,
) (bool, error) {
	return r.queries.CheckUsernameExists(ctx, username)
}

func (r *PostgresUserRepository) Save(ctx context.Context, user *userdomain.User) (*userdomain.User, error) {
	u, err := r.queries.CreateUser(ctx, db.CreateUserParams{
		Username:       user.Username(),
		Email:          user.Email().String(),
		HashedPassword: user.HashedPassword(),
		DisplayName:    user.DisplayName(),
		BirthDay:       user.BirthDay(),
		Gender:         db.GenderType(user.Gender()),
	})
	if err != nil {
		return nil, err
	}

	createdUser := userdomain.NewUserFromDB(u.ID, u.Username, userdomain.Email(u.Email), u.HashedPassword, u.DisplayName, u.BirthDay, userdomain.Gender(u.Gender), u.CreatedAt, u.UpdatedAt)

	return createdUser, nil
}
