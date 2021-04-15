package storages

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages/entities"
)

//go:generate mockgen -source=storage.go -destination=./mocks/storage_mock.go

type StorageManager interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*entities.Task, error)
	AddTask(ctx context.Context, t *entities.Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
	AddUser(ctx context.Context, userID, pwd string) error
}
