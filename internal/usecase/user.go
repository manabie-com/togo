package usecase

import (
	"context"
	"togo/internal/domain"
)

type UserUsecase interface {
	CreateUser(ctx context.Context, user *domain.User) (*domain.User, error)
	GetUserByID(ctx context.Context, id uint) (*domain.User, error)
	UpdateByID(ctx context.Context, id uint, update *domain.User) (*domain.User, error)
}