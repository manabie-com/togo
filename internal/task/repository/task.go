package repository

import (
	"time"
	"togo/internal/models"
	"togo/utils"

	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task *models.Task) (*models.Task, error)
	GetListByUserID(userID int) ([]*models.Task, error)
	GetByID(taskID int) (*models.Task, error)
	GetNumberOfUserTaskOnToday(userID int) (int, error)
}

type taskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

func (t *taskRepository) Create(task *models.Task) (*models.Task, error) {
	if err := t.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (t *taskRepository) GetListByUserID(userID int) ([]*models.Task, error) {
	result := []*models.Task{}
	if err := t.db.Model(&models.Task{}).
		Where("user_id = ?", userID).
		Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (t *taskRepository) GetByID(taskID int) (*models.Task, error) {
	var result models.Task
	if err := t.db.Model(&models.Task{}).
		Where("id = ?", taskID).
		First(&result).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (t *taskRepository) GetNumberOfUserTaskOnToday(userID int) (int, error) {
	var count int64
	now := time.Now()
	startOfday := utils.StartOfDay(now)
	endOfDay := utils.EndOfDay(now)
	if err := t.db.Model(&models.Task{}).
		Where("user_id = ? AND created_at BETWEEN ? AND ? ", userID, startOfday, endOfDay).
		Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}
