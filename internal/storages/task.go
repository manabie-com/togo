package storages

import (
	"context"
)

type TaskStore interface {
	GetTasks(ctx context.Context, task *Task) ([]*Task, error)
	AddTask(ctx context.Context, task *Task) error
}

type TaskCountStore interface {
	Inc(key string) int
	Desc(key string)
	Value(key string) int
}
