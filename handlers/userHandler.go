package handlers

import (
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/manabie-com/togo/database"

	"github.com/manabie-com/togo/models"
)

func GetUserById(togo *models.Togo) (*models.User, error) {

	var user models.User

	if errDb := database.DB.Preload("Tasks").Model(user).Where("Id = ?", togo.Userid).First(&user).Error; errDb != nil {
		if errors.Is(errDb, gorm.ErrRecordNotFound) {
			return CreateUser(togo)
		}
		return nil, errDb
	}

	return &user, nil
}

func CreateUser(togo *models.Togo) (*models.User, error) {

	newUser := &models.User{Id: togo.Userid, LimitTasks: 10, Tasks: []models.Togo{}}
	if errorCreate := database.DB.Create(newUser).Error; errorCreate != nil {
		return nil, errorCreate
	}

	return newUser, nil
}

func deleteUser(togo *models.Togo) error {

	var user models.User

	user.Id = togo.Userid

	if err := database.DB.Model(&user).Where("Id = ?", togo.Userid).Delete(user).Error; err != nil {
		return errors.New("user not found with id: " + strconv.Itoa(togo.Userid))
	}

	return nil

}
