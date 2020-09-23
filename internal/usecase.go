package internal

import (
	"context"
	"database/sql"
	"github.com/manabie-com/togo/internal/storages"
)

type Usecase interface {
	AddTask(ctx context.Context, t *storages.Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error)
}
