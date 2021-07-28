package task

import (
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/task"
)

type taskUsecase struct {
	taskStorage task.TaskStorageInterface
}

func NewTaskUsecase(taskStorage task.TaskStorageInterface) TaskUsecaseInterface {
	return &taskUsecase{
		taskStorage: taskStorage,
	}
}

func (u *taskUsecase) RetrieveTasks(userID, createdDate string) ([]*storages.Task, error) {
	return u.taskStorage.RetrieveTasks(userID, createdDate)
}
