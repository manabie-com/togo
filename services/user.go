package services

import (
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/repositories"
)

type IUserService interface {
	GetUserService(id int) (*models.User, error)
}

type UserService struct {
	repositories.IUserRepository
}

func (userService *UserService) GetUserService(username string) (*models.User, error) {
	return userService.GetUserByUserName(username)
}
