package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/models"
)

type IUserRepository interface {
	GetUserByUserName(username string) (*models.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func (userRepository *UserRepository) GetUserByUserName(username string) (*models.User, error) {
	var user models.User
	result := userRepository.DB.First(&user, username)
	return &user, result.Error
}
