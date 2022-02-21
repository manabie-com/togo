package service

import (
	"context"
	"github.com/google/uuid"

	"github.com/chi07/todo/internal/model"
)

type CreateTaskRepo interface {
	CountUserTasks(ctx context.Context, userID uuid.UUID) (int64, error)
	Create(ctx context.Context, task *model.Task) (uuid.UUID, error)
}

type LimitationRepo interface {
	GetByUserID(ctx context.Context, userID uuid.UUID) (*model.Limitation, error)
}
