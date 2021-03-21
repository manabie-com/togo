package services

import (
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/repositories"
)

type IUserService interface {
	GetUserByUserName(username string) (*models.User, error)
}

type UserService struct {
	UserRepo repositories.IUserRepository
}

func NewUserService(userRepository *repositories.IUserRepository) IUserService {
	return &UserService{UserRepo: *userRepository}
}

func (userService *UserService) GetUserByUserName(username string) (*models.User, error) {
	return userService.UserRepo.GetUserByUserName(username)
}
