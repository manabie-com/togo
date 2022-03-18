package repository

import (
	"context"
	"time"

	"github.com/manabie-com/togo/pkg/errorx"

	"github.com/manabie-com/togo/model"

	"gorm.io/gorm"
)

type TaskRepository interface {
	GetTasks(context.Context, *GetTasksQuery) ([]*model.Task, error)
	GetTask(context.Context, *model.Task) (*model.Task, error)
	SaveTask(*gorm.DB, *model.Task) error
	UpdateTask(*gorm.DB, *model.Task) error
	DeleteTask(*gorm.DB, *model.Task) error
	CountByUserID(context.Context, int) (int64, error)
}

type taskRepository struct {
	*gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) GetTasks(ctx context.Context, query *GetTasksQuery) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := t.Where("user_id = ?", query.UserID).Limit(query.Limit).Offset(query.Offset).Find(&tasks).Error; err != nil {
		return nil, errorx.ErrDatabase(err)
	}
	return tasks, nil
}

func (t *taskRepository) GetTask(ctx context.Context, o *model.Task) (*model.Task, error) {
	task := &model.Task{}
	if err := t.Where(o).
		First(task).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.ErrTaskNotFound(err)
		}
		return nil, errorx.ErrDatabase(err)
	}
	return task, nil
}

func (t *taskRepository) CountByUserID(ctx context.Context, userID int) (int64, error) {
	var count int64
	if err := t.Model(&model.Task{}).Where("user_id = ? and CAST(created_at AS DATE) = ?", userID, time.Now().Format("2006-01-02")).Count(&count).Error; err != nil {
		return 0, errorx.ErrDatabase(err)
	}
	return count, nil
}

func (t *taskRepository) DeleteTask(tx *gorm.DB, task *model.Task) error {
	if err := tx.Delete(task).Error; err != nil {
		return errorx.ErrDatabase(err)
	}
	return nil
}

func (t *taskRepository) SaveTask(tx *gorm.DB, task *model.Task) error {
	if err := tx.Save(task).Error; err != nil {
		return errorx.ErrDatabase(err)
	}
	return nil
}

func (t *taskRepository) UpdateTask(tx *gorm.DB, task *model.Task) error {
	if err := tx.Updates(task).Error; err != nil {
		return errorx.ErrDatabase(err)
	}
	return nil
}
