package handlers

import (
	"errors"
	"strconv"
	"time"

	"github.com/manabie-com/togo/database"
	"github.com/manabie-com/togo/models"
)

func AddTodo(newTodo *models.Togo) (*models.User, error) {

	var user models.User

	if err := database.DB.Preload("Tasks").Where("Id = ?", newTodo.Userid).First(&user).Error; err != nil {
		return nil, errors.New("user not found with id: " + strconv.Itoa(newTodo.Userid))
	}

	if user.CountTasks() >= 10 {
		return nil, errors.New("tasks has been limit." + strconv.Itoa(user.CountTasks()))
	}
	// Create todo task
	todo := models.Togo{Task: newTodo.Task, Userid: newTodo.Userid, Date: time.Now()}

	database.DB.Model(todo).Create(&todo)

	user.Tasks = append(user.Tasks, todo)

	return &user, nil

}

func resetLimitTask(todo *models.Togo) (*models.User, error) {

	var user models.User

	if err := database.DB.Preload("Tasks").Where("Id = ?", todo.Userid).First(&user).Error; err != nil {
		return nil, errors.New("user not found with id: " + strconv.Itoa(todo.Userid))
	}

	// update countTasks
	user.Tasks = []models.Togo{}
	database.DB.Model(&user).Update(user)

	return &user, nil
}

func deleteTodo(todo *models.Togo) error {

	if err := database.DB.Model(&todo).Where("userid = ?", todo.Userid).Delete(todo).Error; err != nil {
		return errors.New("user not found with id: " + strconv.Itoa(todo.Userid))
	}

	return nil

}
