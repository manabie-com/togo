package storages

import (
	"context"
	"github.com/manabie-com/togo/internal/entities"
)

var userRepository *UserRepository
var taskRepository *TaskRepository

type UserRepository interface {
	ValidateUser(ctx context.Context, userID, pwd string) bool
}

type TaskRepository interface {
	RetrieveTasks(ctx context.Context, userID, createdDate string) ([]*entities.Task, error)
	AddTask(ctx context.Context, t *entities.Task) error
	CountTaskPerDayByUserID(ctx context.Context, userID string) (uint, error)
}

func SetUserRepository(repository UserRepository) {
	userRepository = &repository
}

func GetUserRepository() UserRepository {
	return *userRepository
}

func SetTaskRepository(repository TaskRepository) {
	taskRepository = &repository
}

func GetTaskRepository() TaskRepository {
	return *taskRepository
}
