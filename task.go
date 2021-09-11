package togo

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/storages"
)

type TaskService interface {
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error)
	AddTask(ctx context.Context, t *storages.Task) error
}
