package repository

import (
	entity "github.com/manabie-com/togo/internal/entities"

	postgre "github.com/manabie-com/togo/internal/storages/postgre"
)

// TaskRepository action CRUD with Task entity
type TaskRepository struct {
	Store *postgre.Storage
}

// Add func add new task into Db
func (repo *TaskRepository) Add(entity *entity.Task) (*entity.Task, error) {

	result, err := repo.Store.Add(entity)

	return result, err
}

// GetAll func retrives all task in Db
func (repo *TaskRepository) GetAll(createdDate string) ([]entity.Task, error) {

	result, err := repo.Store.GetAll(createdDate)

	return result, err
}
