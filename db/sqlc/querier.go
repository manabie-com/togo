// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"context"
)

type Querier interface {
	CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	GetTask(ctx context.Context, id int64) (Task, error)
	GetUser(ctx context.Context, username string) (User, error)
	ListTasksByOwner(ctx context.Context, arg ListTasksByOwnerParams) ([]Task, error)
}

var _ Querier = (*Queries)(nil)
