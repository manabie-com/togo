package storages

import (
	"context"
)

//go:generate mockgen -package mock -destination mock/store_mock.go . Store
type Store interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string, page, limit int) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
}
