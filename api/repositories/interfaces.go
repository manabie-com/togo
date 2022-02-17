package repositories

import "github.com/kier1021/togo/api/models"

type IUserTaskRepository interface {
	CreateUser(user models.User) (string, error)
	AddTaskToUser(user models.User, userTask models.Task, insDay string) error
	GetUser(filter map[string]interface{}) (*models.User, error)
	GetUserTask(filter map[string]interface{}) (*models.UserTask, error)
}
