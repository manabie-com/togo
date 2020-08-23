package storages

import (
	"context"
	"database/sql"
)

// Store - store interface
type Store interface {
	AddTask(ctx context.Context, t *Task) error
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	ValidateUser(ctx context.Context, userID sql.NullString, pwd string) bool
}
