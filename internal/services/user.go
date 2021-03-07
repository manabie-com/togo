package services

import (
	"context"
	"github.com/manabie-com/togo/internal/model"
)

type UserService interface {
	GetAuthToken(ctx context.Context, id, password string) (bool, error)
}

type userService struct {
	userStorage model.UserStorage
}
func NewUserService(us model.UserStorage) UserService{
	return &userService{
		userStorage: us,
	}
}

func (us *userService) GetAuthToken(ctx context.Context, id, password string) (bool, error) {
	isValidUser, err := us.userStorage.ValidateUser(ctx, id, password)
	return isValidUser, err
}

