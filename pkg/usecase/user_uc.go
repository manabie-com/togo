package usecase

import (
	"context"
	"togo.com/pkg/model"
	"togo.com/pkg/repository"
)

type UserUseCase interface {
	Login(ctx context.Context, request model.LoginRequest) error
}
type userUseCase struct {
	repo repository.Repository
}

func NewUserUseCase(repo repository.Repository) UserUseCase {
	return userUseCase{repo: repo}
}

func (t userUseCase) Login(ctx context.Context, request model.LoginRequest) error {
	return nil
}
