package repositories

import (
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/utils"
	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserByUserName(username string) (*models.User, error)
	AddUser(user *models.User) (*models.User, error)
}

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{DB: db}
}

func (userRepo *UserRepository) GetUserByUserName(username string) (*models.User, error) {
	var user models.User
	result := userRepo.DB.First(&user, "username = ?", username)
	return &user, result.Error
}

func (userRepo *UserRepository) AddUser(user *models.User) (*models.User, error) {
	hashedPass, err := utils.Hash(user.Password)

	if err != nil {
		return user, err
	}

	user.Password = hashedPass

	result := userRepo.DB.Create(user)
	return user, result.Error
}
