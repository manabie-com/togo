package services

import (
	"github.com/manabie-com/togo/entities"
	"github.com/manabie-com/togo/helpers"
	"github.com/manabie-com/togo/repositories"
)

type ITaskService interface {
	GetAllTask() ([]entities.Task, *helpers.Response)
}

type TaskService struct {
	Repo repositories.ITaskRepository
}

func NewTaskService(repo repositories.ITaskRepository) ITaskService {
	return &TaskService{
		Repo: repo,
	}
}

func (taskService *TaskService) GetAllTask() ([]entities.Task, *helpers.Response) {
	tasks, resErr := taskService.Repo.GetAllTask()
	if resErr != nil {
		return nil, resErr
	}

	return tasks, nil
}
