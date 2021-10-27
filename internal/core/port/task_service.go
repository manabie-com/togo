package port

import (
	"context"

	"github.com/manabie-com/togo/internal/core/domain"
)

type TaskService interface {
	RetrieveTasks(ctx context.Context, userId, createdDate string) ([]*domain.Task, error)
	AddTask(ctx context.Context, task *domain.Task) error
	Login(ctx context.Context, username, password string) (string, error)
}
