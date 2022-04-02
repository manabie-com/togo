package repository

import (
	"context"
	"togo/internal/domain"
)

// UserRepository repository interface
type UserRepository interface {
	Create(ctx context.Context, entity *domain.User) (*domain.User, error)
	FindOne(ctx context.Context, filter *domain.User) (*domain.User, error)
	UpdateByID(ctx context.Context, id uint, update *domain.User) (*domain.User, error)
}
