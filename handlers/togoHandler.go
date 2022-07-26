package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/manabie-com/togo/database"
	"github.com/manabie-com/togo/models"
)

func Addtogo(newtogo *models.Togo) (*models.User, error) {

	var user models.User

	if err := database.DB.Preload("Tasks").Where("Id = ?", newtogo.Userid).First(&user).Error; err != nil {
		return nil, errors.New("user not found with id: " + strconv.Itoa(newtogo.Userid))
	}

	if user.CountTasks() >= 10 {
		return nil, errors.New("tasks has been limit." + strconv.Itoa(user.CountTasks()))
	}
	// Create togo task
	togo := models.Togo{Task: newtogo.Task, Userid: newtogo.Userid, Date: time.Now()}

	database.DB.Model(togo).Create(&togo)

	user.Tasks = append(user.Tasks, togo)

	return &user, nil

}

func resetLimitTask(togo *models.Togo) (*models.User, error) {

	var user models.User

	if err := database.DB.Preload("Tasks").Where("Id = ?", togo.Userid).First(&user).Error; err != nil {
		return nil, errors.New("user not found with id: " + strconv.Itoa(togo.Userid))
	}

	// update countTasks
	user.Tasks = []models.Togo{}
	database.DB.Model(&user).Update(user)

	return &user, nil
}

func deletetogo(togo *models.Togo) error {

	if err := database.DB.Model(&togo).Where("userid = ?", togo.Userid).Delete(togo).Error; err != nil {
		return errors.New("user not found with id: " + strconv.Itoa(togo.Userid))
	}

	return nil

}
