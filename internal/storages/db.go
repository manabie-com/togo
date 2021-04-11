package storages

import (
	"context"
	"database/sql"
	"errors"
)

var DailyLimitExceededError = errors.New("Daily tasks limit exceeded")

// DB stores tasks and credentials of users
type DB interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*Task, error)
	AddTask(ctx context.Context, t *Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
}
