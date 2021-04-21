package usecase

import (
	"context"

	"github.com/manabie-com/togo/internal/model"
)

type UserService interface {
	GetAuthToken(ctx context.Context, id, password string) (bool, error)
}

type userService struct {
	userRespository model.UserRespository
}

func NewUserService(ur model.UserRespository) UserService {
	return &userService{
		userRespository: ur,
	}
}

func (us *userService) GetAuthToken(ctx context.Context, id, password string) (bool, error) {
	isValidUser, err := us.userRespository.ValidateUser(ctx, id, password)
	return isValidUser, err
}
