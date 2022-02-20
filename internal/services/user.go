package services

import (
	"context"
	"fmt"
	"togo/internal/domain"
	"togo/internal/provider"
	"togo/internal/repository"
)

type userService struct {
	passwordHashProvider provider.PasswordHashProvider
	userRepo             repository.UserRepository
}

// NewUserService service constructor
func NewUserService(
	passwordHashProvider provider.PasswordHashProvider,
	userRepo repository.UserRepository,
) domain.UserService {
	return &userService{
		passwordHashProvider: passwordHashProvider,
		userRepo:             userRepo,
	}
}

func (s userService) CreateUser(ctx context.Context, input *domain.User) (*domain.User, error) {
	// Existing check
	userExists, err := s.userRepo.FindOne(ctx, &domain.User{Username: input.Username})
	if err != nil && err != domain.ErrUserNotFound {
		return nil, fmt.Errorf("userService:CreateUser: %w", err)
	}
	if userExists != nil {
		return nil, domain.ErrDuplicatedUsername
	}
	// Hash the password
	pwd, err := s.passwordHashProvider.HashPassword(input.Password)
	if err != nil {
		return nil, fmt.Errorf("userService:CreateUser: %w", err)
	}
	input.Password = pwd
	user, err := s.userRepo.Create(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("userService:CreateUser: %w", err)
	}
	return user, nil
}

func (s userService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	user, err := s.userRepo.FindOne(ctx, &domain.User{ID: id})
	if err != nil {
		return nil, fmt.Errorf("userService:GetUserByID: %w", err)
	}
	return user, nil
}

func (s userService) UpdateByID(ctx context.Context, id uint, update *domain.User) (*domain.User, error) {
	user, err := s.userRepo.UpdateByID(ctx, id, update)
	if err != nil {
		return nil, fmt.Errorf("userService:UpdateByID: %w", err)
	}
	return user, nil
}
