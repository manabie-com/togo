package repositories

import "github.com/kier1021/togo/api/models"

type IUserTaskRepository interface {
	AddTaskToUser(user models.User, userTask models.Task, insDay string) error
	GetUserTask(filter map[string]interface{}) (*models.UserTask, error)
}

type IUserRepository interface {
	CreateUser(user models.User) (string, error)
	GetUser(filter map[string]interface{}) (*models.User, error)
	GetUsers(filter map[string]interface{}) ([]models.User, error)
}
