package storages

import (
	"context"
)

type Store interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string, page, limit int) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
}
