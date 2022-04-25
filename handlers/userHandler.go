package handlers

import (
	"errors"
	"strconv"

	"github.com/jinzhu/gorm"

	"github.com/manabie-com/togo/database"

	"github.com/manabie-com/togo/models"
)

func GetUserById(todo *models.Todo) (*models.User, error) {

	var user models.User

	if errDb := database.DB.Preload("Tasks").Model(user).Where("Id = ?", todo.Userid).First(&user).Error; errDb != nil {
		if errors.Is(errDb, gorm.ErrRecordNotFound) {
			return CreateUser(todo)
		}
		return nil, errDb
	}

	return &user, nil
}

func CreateUser(todo *models.Todo) (*models.User, error) {

	newUser := &models.User{Id: todo.Userid, LimitTasks: 10, Tasks: []models.Todo{}}
	if errorCreate := database.DB.Create(newUser).Error; errorCreate != nil {
		return nil, errorCreate
	}

	return newUser, nil
}

func deleteUser(todo *models.Todo) error {

	var user models.User

	user.Id = todo.Userid

	if err := database.DB.Model(&user).Where("Id = ?", todo.Userid).Delete(user).Error; err != nil {
		return errors.New("user not found with id: " + strconv.Itoa(todo.Userid))
	}

	return nil

}
