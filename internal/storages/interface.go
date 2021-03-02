package storages

import (
	"context"
	"database/sql"
)

type StorageDB interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}