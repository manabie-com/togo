// Code generated by sqlc. DO NOT EDIT.

package postgres

import (
	"context"
)

type Querier interface {
	CreateTask(ctx context.Context, arg CreateTaskParams) (Task, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	RetrieveTasks(ctx context.Context, arg RetrieveTasksParams) ([]Task, error)
	RetrieveUser(ctx context.Context, username string) (User, error)
}

var _ Querier = (*Queries)(nil)
