package usecases

import (
	"context"
	"fmt"
	"togo/internal/pkg/domain/dtos"
	"togo/internal/pkg/repositories"
	"togo/pkg/auth"
	"togo/pkg/utils"
)

type AuthUsecase interface {
	Login(cxt context.Context, req dtos.LoginRequest) (string, error)
}

type authUsecase struct {
	repo repositories.UserRepository
}

func (u *authUsecase) Login(ctx context.Context, req dtos.LoginRequest) (string, error) {
	user, err := u.repo.FindUserWithEmail(ctx, req.Email)

	if err != nil {
		return "", err
	}
	if utils.EncrtyptPasswords(req.Password) != user.Password {
		return "", fmt.Errorf("INVALID_PASSWORD")
	}

	tokenString, err := auth.GenerateJWT(user.ID)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// NewAuthUsecase
func NewAuthUsecase(userRepo repositories.UserRepository) AuthUsecase {
	return &authUsecase{
		repo: userRepo,
	}
}
