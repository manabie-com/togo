package repository

import (
	"context"
	"togo/domain/model"
)

type UserRepository interface {
	Create(ctx context.Context, u model.User) error
	Delete(ctx context.Context, username string) error
	Get(ctx context.Context, username string) (u model.User, err error)
}

type TaskRepository interface {
	CountTaskCreatedInDayByUser(ctx context.Context, u model.User) (int, error)
	Create(ctx context.Context, u model.Task) error
}

type TokenResponse interface {
}
