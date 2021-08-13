package repository

import "github.com/manabie-com/togo/internal/model"

type TasksRepository interface {
	GetByIdAndCreateDate(id string, createdDate string) (model.TaskList, error)
	Save(task *model.Task) (*model.Task, error)
	CountByIdAndCreateDate(id string, createdDate string) (int, error)
}
