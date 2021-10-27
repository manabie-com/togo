package port

import (
	"context"

	"github.com/manabie-com/togo/internal/core/domain"
	"github.com/manabie-com/togo/pkg/database"
)

type TaskRepository interface {
	// RetrieveTasks returns tasks if match userId AND createDate.
	RetrieveTasks(ctx context.Context, conn database.Connection, userId, createdDate string) ([]*domain.Task, error)

	// AddTask adds a new task to DB
	AddTask(ctx context.Context, conn database.Connection, task *domain.Task) error
	Login(ctx context.Context, conn database.Connection, username, password string) (string, error)
}
