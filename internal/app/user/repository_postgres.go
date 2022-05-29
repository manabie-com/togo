package user

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/dinhquockhanh/togo/internal/pkg/errors"
	db "github.com/dinhquockhanh/togo/internal/pkg/sql/sqlc"
	"github.com/dinhquockhanh/togo/internal/pkg/util"
)

type (
	PostgresRepository struct {
		q *db.Queries
	}
)

func NewPostgresRepository(cnn *sql.DB) Repository {
	return &PostgresRepository{
		q: db.New(cnn),
	}
}

func (r *PostgresRepository) Create(ctx context.Context, req *CreateUserReq) (*UserSafe, error) {
	op := "postgresRepository.Create"

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)

	}

	user, err := r.q.CreateUser(ctx, &db.CreateUserParams{
		Username:       req.UserName,
		FullName:       req.FullName,
		HashedPassword: hashedPassword,
		Email:          req.Email,
		CreatedAt:      time.Now(),
		TierID:         1, // TODO: check tier id in db first
	})

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &UserSafe{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		TierID:    user.TierID,
	}, nil
}

func (r *PostgresRepository) UpdateUserTier(ctx context.Context, req *UpdateUserTierReq) (*UserSafe, error) {
	user, err := r.q.UpdateUserTier(ctx, &db.UpdateUserTierParams{
		TierID:   req.TierID,
		Username: req.UserName,
	})

	if err != nil {
		op := "postgresRepository.UpdateUserTier"
		if errors.IsSQLNotFound(err) {
			return nil, errors.NewNotFoundErr(err, op, fmt.Sprintf("user with username = %s", req.UserName))
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &UserSafe{
		Username:  user.Username,
		FullName:  user.FullName,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		TierID:    user.TierID,
	}, nil
}

func (r *PostgresRepository) GetByUserName(ctx context.Context, req *GetUserByUserNameReq) (*User, error) {
	user, err := r.q.GetUserByName(ctx, req.UserName)
	if err != nil {
		op := "postgresRepository.GetByUserName"
		if errors.IsSQLNotFound(err) {
			return nil, errors.NewNotFoundErr(err, op, fmt.Sprintf("user with username = %s", req.UserName))
		}
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	return &User{
		Username:       user.Username,
		FullName:       user.FullName,
		HashedPassword: user.HashedPassword,
		Email:          user.Email,
		CreatedAt:      user.CreatedAt,
		TierID:         user.TierID,
	}, nil

}

func (r *PostgresRepository) List(ctx context.Context, req *ListUsersReq) ([]*UserSafe, error) {
	//TODO implement me
	panic("implement me")
}

func (r *PostgresRepository) Delete(ctx context.Context, req *DeleteUserByNameReq) error {
	err := r.q.DeleteUser(ctx, req.UserName)
	if err != nil {
		op := "postgresRepository.Delete"
		if errors.IsSQLNotFound(err) {
			return errors.NewNotFoundErr(err, op, fmt.Sprintf("user with username = %s", req.UserName))
		}
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
