package repository

import (
	"context"
	"togo/internal/domain"
)

// TaskRepository repository interface
type TaskRepository interface {
	Create(ctx context.Context, entity *domain.Task) (*domain.Task, error)
	Update(ctx context.Context, filter, update *domain.Task) (*domain.Task, error)
	Find(ctx context.Context, filter *domain.Task) ([]*domain.Task, error)
}
