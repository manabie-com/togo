package repositories

import (
	"github.com/manabie-com/togo/models"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByUserName(username string) (*models.User, error)
	AddUser(user *models.User) (*models.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (userRepository *UserRepository) GetUserByUserName(username string) (*models.User, error) {
	var user models.User
	result := userRepository.DB.First(&user, username)
	return &user, result.Error
}

func (userRepository *UserRepository) AddUser(user *models.User) (*models.User, error) {
	result := userRepository.DB.Create(user)
	return user, result.Error
}
