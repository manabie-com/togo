package storages

import (
	"context"
)

type TaskStore interface {
	RetrieveTasks(ctx context.Context, task *Task) ([]*Task, error)
	AddTask(ctx context.Context, task *Task) error
}

type TaskCountStore interface {
	CreateIfNotExists(ctx context.Context, userID, date string) error
	UpdateAndGet(ctx context.Context, userID, date string) (int, error)
}
