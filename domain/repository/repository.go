package repository

import (
	"context"
	"togo/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, u model.User) error
	Get(ctx context.Context, username string) (u model.User, err error)
}

type TaskRepository interface {
	Create(ctx context.Context, u model.Task) error
}

type TokenResponse interface {
}
