package repository

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"
	"github.com/trinhdaiphuc/togo/database/ent/user"
	"github.com/trinhdaiphuc/togo/internal/entities"
	"github.com/trinhdaiphuc/togo/internal/infrastructure"
)

type UserRepository interface {
	GetUserByName(ctx context.Context, username string) (*entities.User, error)
}

type userRepositoryImpl struct {
	db *infrastructure.DB
}

func NewUserRepository(db *infrastructure.DB) UserRepository {
	return &userRepositoryImpl{
		db: db,
	}
}

func (u *userRepositoryImpl) GetUserByName(ctx context.Context, username string) (*entities.User, error) {
	resp, err := u.db.User.Query().Where(user.UsernameEQ(username)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return &entities.User{
		ID:        resp.ID,
		Username:  resp.Username,
		Password:  resp.Password,
		TaskLimit: resp.TaskLimit,
	}, nil
}
