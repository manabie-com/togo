package usecase

import (
	"context"
	"togo/internal/domain"
)

type TaskUsecase interface {
	Create(ctx context.Context, task *domain.Task) (*domain.Task, error)
	Update(ctx context.Context, filter, update *domain.Task) (*domain.Task, error)
	FindByUserID(ctx context.Context, userID uint) ([]*domain.Task, error)
}