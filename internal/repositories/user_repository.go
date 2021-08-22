package repositories

import (
	"fmt"
	"gorm.io/gorm"
	models "github.com/manabie-com/togo/internal/models"
)

type UserRepository struct {
	DB *gorm.DB
}

func ProvideUserRepository(DB *gorm.DB) UserRepository {
	return UserRepository{DB: DB}
}

func (repo *UserRepository) Create(user models.User) models.User {
	result := repo.DB.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		panic("Error, insert user to into database!")
	}
	return user
}

func (repo *UserRepository) FindWhere(condition map[string]interface{}) (models.User, error) {
	var user models.User
	result := repo.DB.Where(condition).First(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		fmt.Println("Error, select user by condition!")
	}
	return user, result.Error
}
