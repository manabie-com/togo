package mocks

import (
	"github.com/manabie-com/togo/models"
	"github.com/stretchr/testify/mock"
)

type ITaskRepository struct {
	mock.Mock
}

func (mock *ITaskRepository) GetTasksByUserName(username string, createdAt string) (*[]models.Task, error) {
	args := mock.Called(username, createdAt)
	return args.Get(0).(*[]models.Task), args.Error(1)
}

func (mock *ITaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	args := mock.Called(task)
	return args.Get(0).(*models.Task), args.Error(1)
}

func (mock *ITaskRepository) Count(username string) (int64, error) {
	args := mock.Called(username)
	return args.Get(0).(int64), args.Error(1)
}
