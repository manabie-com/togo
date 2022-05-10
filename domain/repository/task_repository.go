package repository

import "github.com/jfzam/togo/domain/entity"

type TaskRepository interface {
	SaveTask(*entity.Task, int64) (*entity.Task, map[string]string)
	GetTask(uint64) (*entity.Task, error)
	GetAllTask() ([]entity.Task, error)
}
