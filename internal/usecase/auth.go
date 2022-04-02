package usecase

import (
	"context"
	"togo/internal/domain"
)

type AuthUsecase interface {
	Login(ctx context.Context, credential *domain.LoginCredential) (*domain.LoginResult, error)
	VerifyToken(ctx context.Context, token string) (*domain.VerifyTokenResult, error)
}