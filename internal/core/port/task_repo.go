package port

import (
	"context"

	"github.com/manabie-com/togo/internal/core/domain"
	"github.com/manabie-com/togo/pkg/database"
)

type TaskRepository interface {
	InitTables(ctx context.Context, conn database.Connection) error
	RetrieveTasks(ctx context.Context, conn database.Connection, userId, createdDate string) ([]*domain.Task, error)
	AddTask(ctx context.Context, conn database.Connection, task *domain.Task) error
	Login(ctx context.Context, conn database.Connection, username, password string) (string, error)
}
