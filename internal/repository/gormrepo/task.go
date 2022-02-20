package gormrepo

import (
	"context"
	"errors"
	"fmt"
	"togo/internal/domain"
	"togo/internal/repository"

	"gorm.io/gorm"
)

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository repository constructor
func NewTaskRepository(db *gorm.DB) repository.TaskRepository {
	db.AutoMigrate(&domain.Task{})
	return &taskRepository{
		db,
	}
}

func (r taskRepository) Create(ctx context.Context, entity *domain.Task) (*domain.Task, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, fmt.Errorf("taskRepository:Created: %w", err)
	}
	return entity, nil
}

func (r taskRepository) UpdateByID(ctx context.Context, id uint, update *domain.Task) (*domain.Task, error) {
	task := &domain.Task{ID: id}
	if err := r.db.Model(task).Updates(update).Error; err != nil {
		return nil, fmt.Errorf("taskRepository:UpdateByID: %w", err)
	}
	return task, nil
}

func (r taskRepository) Find(ctx context.Context, filter *domain.Task) ([]*domain.Task, error) {
	tasks := make([]*domain.Task, 0)
	if err := r.db.Where(filter).Find(&tasks).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return tasks, fmt.Errorf("taskRepository:Find: %w", err)
	}
	return tasks, nil
}
