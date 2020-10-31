package storages

import (
	"context"
	"database/sql"

	"github.com/manabie-com/togo/internal/entities"
)

// Storages interface (like repository layer)
type Storages interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]entities.Task, error)
	AddTask(ctx context.Context, t entities.Task) error
	ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
	GetMaxTaskTodo(ctx context.Context, userID string) (int, error)
}
