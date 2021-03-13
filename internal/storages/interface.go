package storages

import (
	"context"
	"database/sql"
)

type StoreInterface interface {
// RetrieveTasks returns tasks if match userID AND createDate.
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)

	// AddTask adds a new task to DB
	AddTask(ctx context.Context, t *Task) error

	// ValidateUser returns tasks if match userID AND password
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}