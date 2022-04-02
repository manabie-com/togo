package services

import (
	"context"
	"fmt"
	"togo/internal/domain"
	"togo/internal/provider"
	"togo/internal/repository"
)

type authService struct {
	passwordHashProvider provider.PasswordHashProvider
	tokenProvider        provider.TokenProvider

	userRepo repository.UserRepository
}

// NewAuthService service constructor
func NewAuthService(
	passwordHashProvider provider.PasswordHashProvider,
	tokenProvider provider.TokenProvider,
	userRepo repository.UserRepository,
) domain.AuthService {
	return &authService{
		passwordHashProvider: passwordHashProvider,
		tokenProvider:        tokenProvider,
		userRepo:             userRepo,
	}
}

func (s authService) Login(ctx context.Context, credential *domain.LoginCredential) (*domain.LoginResult, error) {
	user, err := s.userRepo.FindOne(ctx, &domain.User{Username: credential.Username})
	if err != nil {
		return nil, domain.ErrCredentialInvalid
	}
	err = s.passwordHashProvider.ComparePassword(credential.Password, user.Password)
	if err != nil {
		return nil, domain.ErrCredentialInvalid
	}
	token, err := s.tokenProvider.GenerateToken(user)
	if err != nil {
		return nil, domain.ErrLoginFailed
	}
	return &domain.LoginResult{
		Profile: user,
		Token:   token,
	}, nil
}

func (s authService) VerifyToken(ctx context.Context, token string) (*domain.VerifyTokenResult, error) {
	payload, err := s.tokenProvider.VerifyToken(token)
	if err != nil {
		return nil, fmt.Errorf("authService:VerifyToken: %w", err)
	}
	user, err := s.userRepo.FindOne(ctx, &domain.User{ID: payload.(*domain.User).ID})
	if err != nil {
		return nil, domain.ErrUserNotFound
	}
	return &domain.VerifyTokenResult{
		Authenticated: true,
		Payload:       user,
	}, nil
}
