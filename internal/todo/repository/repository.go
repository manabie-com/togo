package repository

import (
	"context"

	"github.com/manabie-com/togo/pkg/errorx"

	"github.com/manabie-com/togo/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(context.Context, GetTasksQuery) ([]*model.Task, error)
	CreateTask(context.Context, *model.Task) error
	CountByUserID(context.Context, int) (int64, error)
}

type taskRepository struct {
	*gorm.DB
}

func (t *taskRepository) CreateTask(ctx context.Context, task *model.Task) error {
	if err := t.Save(task).Error; err != nil {
		return errorx.ErrDatabase(err)
	}
	return nil
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) GetTasks(ctx context.Context, query GetTasksQuery) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := t.Where("user_id = ?", query.UserID).Find(&tasks).Error; err != nil {
		return nil, errorx.ErrDatabase(err)
	}
	return tasks, nil
}

func (t *taskRepository) CountByUserID(ctx context.Context, userID int) (int64, error) {
	var count int64
	if err := t.Model(&model.Task{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return 0, errorx.ErrDatabase(err)
	}
	return count, nil
}
