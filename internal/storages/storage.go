package storages

import (
	"context"
	"time"

	"github.com/manabie-com/togo/internal/storages/entities"
)

//go:generate mockgen -source=storage.go -destination=./mocks/storage_mock.go

type StorageManager interface {
	RetrieveTasks(ctx context.Context, userID string, date time.Time) ([]*entities.Task, error)
	AddTask(ctx context.Context, task *entities.Task) error
	ValidateUser(ctx context.Context, userID, pwd string) bool
	AddUser(ctx context.Context, userID, pwd string) error
}
