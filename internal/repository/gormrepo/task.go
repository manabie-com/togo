package gormrepo

import (
	"context"
	"errors"
	"togo/internal/domain"
	"togo/internal/repository"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository repository constructor
func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	return &taskRepository{
		db,
	}
}

func (r taskRepository) Create(ctx context.Context, entity *domain.Task) (*domain.Task, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

func (r taskRepository) UpdateByID(ctx context.Context, id uint, update *domain.Task) (*domain.Task, error) {
	task := new(domain.Task)
	if err := r.db.Model(task).Updates(update).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r taskRepository) Find(ctx context.Context, filter *domain.Task) ([]*domain.Task, error) {
	tasks := make([]*domain.Task, 0)
	if err := r.db.Where(filter).Find(&tasks).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return tasks, err
	}
	return tasks, nil
}
