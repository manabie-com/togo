package repositories

import "github.com/kier1021/togo/api/models"

// IUserTaskRepository is the interface for user task repositories
// Implement this interface to be used in UserTaskService
type IUserTaskRepository interface {
	AddTaskToUser(user models.User, userTask models.Task, insDay string) error
	GetUserTask(filter map[string]interface{}) (*models.UserTask, error)
}

// IUserRepository is the interface for user repositories
// Implement this interface to be used in UserService
type IUserRepository interface {
	CreateUser(user models.User) (string, error)
	GetUser(filter map[string]interface{}) (*models.User, error)
	GetUsers(filter map[string]interface{}) ([]models.User, error)
}
