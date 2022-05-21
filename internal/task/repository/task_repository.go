package repository

import (
	"gorm.io/gorm"
	"togo/domain"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.ITaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (repo *TaskRepository) Create(task domain.Task) (domain.Task, error) {
	result := repo.db.Create(&task)

	if result.Error != nil {
		return domain.Task{}, result.Error
	}
	return task, nil
}

func (repo *TaskRepository) Save(task domain.Task) error {
	result := repo.db.Save(&task)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (repo *TaskRepository) SetTx(tx *gorm.DB) *gorm.DB {
	repo.db = tx
	return tx
}
