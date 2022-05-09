package repository

import "github.com/jfzam/togo/domain/entity"

type FoodRepository interface {
	SaveTask(*entity.Task) (*entity.Task, map[string]string)
	GetTask(uint64) (*entity.Task, error)
	GetAllTask() ([]entity.Task, error)
}
