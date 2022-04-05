package usecases

import (
	"togo/internal/pkg/repositories"
)

type UserUsecase interface {
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

// NewUserUsecase
func NewUserUsecase(userRepo repositories.UserRepository) UserUsecase {
	return &userUsecase{
		userRepo: userRepo,
	}
}
