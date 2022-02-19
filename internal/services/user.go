package services

import (
	"context"
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

// CreateUser method for create new User
func (s userService) CreateUser(ctx context.Context, user *domain.User) (*domain.User, error) {
	// Hash the password
	pwd, err := s.passwordHashProvider.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}
	user.Password = pwd
	return s.userRepo.Create(ctx, user)
}

// GetUserByID method for get one User
func (s userService) GetUserByID(ctx context.Context, id uint) (*domain.User, error) {
	return s.userRepo.FindOne(ctx, &domain.User{ID: id})
}
