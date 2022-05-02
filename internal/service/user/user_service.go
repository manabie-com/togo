package userservice

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

import (
	"context"
	"todo/configs"
	"todo/internal/entities"
	"todo/internal/repository"
)

type UserService interface {
	Login(ctx context.Context, user *entities.User) (*entities.User, error)
	SignUp(ctx context.Context, user *entities.User) (*entities.User, error)
}

type userService struct {
	userRepo repository.UserRepository
	cfg      *configs.Config
}

func NewUserService(userRepo repository.UserRepository, cfg *configs.Config) UserService {
	return &userService{
		userRepo: userRepo,
		cfg:      cfg,
	}
}
