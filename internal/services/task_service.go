package services

import "github.com/manabie-com/togo/internal/model"

type TasksService interface {
	RetrieveTasks(id string, createdDate string) (model.TaskList, error)
	AddTask(task *model.Task, userID string) (*model.Task, error)
}

