package storages

import (
	"context"
	"database/sql"
)

type DB interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
	RetrieveUser(ctx context.Context, userID string) (*User, error)
	CountTasks(ctx context.Context, userID, createdDate string) (count sql.NullInt32, err error)
}
