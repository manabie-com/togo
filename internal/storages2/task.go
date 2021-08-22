package storages2

import (
	"context"
	"github.com/jmoiron/sqlx"
)

type ITaskRepo interface {
	RetrieveTasks(ctx context.Context, arg RetrieveTasksParams) ([]Task, error)
	AddTask(ctx context.Context, arg AddTaskParams) error
}

type TaskRepo struct {
	*Queries
	db *sqlx.DB
}

func NewTaskRepo(db *sqlx.DB) ITaskRepo {
	return &TaskRepo{
		db: db,
	}
}
