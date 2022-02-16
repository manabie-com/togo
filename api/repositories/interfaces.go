package repositories

import "github.com/kier1021/togo/api/models"

type IUserTaskRepository interface {
	CreateUser(user models.User) (string, error)
	AddTaskToUser(userTask models.UserTask) error
	GetUser(filter map[string]interface{}) (*models.User, error)
}
