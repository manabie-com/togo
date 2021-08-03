package handler

import (
	"context"

	"togo/common"
	"togo/internal/entity"
)

type UserRepository interface {
	CreateUser(ctx context.Context, username string, password string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUser(ctx context.Context, id int32) (*entity.User, error)
}

type UserRedisRepo interface {
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetUser(ctx context.Context, id int32) (*entity.User, error)
	SetUser(ctx context.Context, user *entity.User) error
}

type UserHandler struct {
	repo    UserRepository
	rdbRepo UserRedisRepo
}

func NewUserHandler(repo UserRepository, rdbRepo UserRedisRepo) *UserHandler {
	return &UserHandler{repo: repo, rdbRepo: rdbRepo}
}

func (u *UserHandler) CreateUser(ctx context.Context, username string, password string) (*entity.User, error) {
	userRdb, err := u.rdbRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if userRdb != nil {
		return nil, common.ErrUserAlreadyExist
	}

	if _, err = u.repo.GetUserByUsername(ctx, username); err == nil {
		return nil, common.ErrUserAlreadyExist
	}

	user, err := u.repo.CreateUser(ctx, username, password)
	if err != nil {
		return nil, err
	}

	_ = u.rdbRepo.SetUser(ctx, user)

	return user, nil
}

func (u *UserHandler) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {
	userRdb, err := u.rdbRepo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	if userRdb != nil {
		return userRdb, nil
	}

	user, err := u.repo.GetUserByUsername(ctx, username)
	if err != nil {
		return nil, common.ErrUserNotFound
	}

	_ = u.rdbRepo.SetUser(ctx, user)

	return user, nil
}

func (u *UserHandler) GetUser(ctx context.Context, id int32) (*entity.User, error) {
	userRdb, err := u.rdbRepo.GetUser(ctx, id)
	if err != nil {
		return nil, err
	}

	if userRdb != nil {
		return userRdb, nil
	}

	user, err := u.repo.GetUser(ctx, id)
	if err != nil {
		return nil, common.ErrUserNotFound
	}

	_ = u.rdbRepo.SetUser(ctx, user)

	return user, nil
}
