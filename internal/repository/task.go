package repository

import (
	"time"
	"togo/internal/models"
	"togo/utils"

	"gorm.io/gorm"
)

type TaskRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{db}
}

func (t *TaskRepository) WithTrx(trxHandle *gorm.DB) *TaskRepository {
	if trxHandle == nil {
		return t
	}
	t.db = trxHandle
	return t
}

func (t *TaskRepository) Create(task *models.Task) (*models.Task, error) {
	if err := t.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskRepository) GetListByUserID(userID int) ([]*models.Task, error) {
	result := []*models.Task{}
	if err := t.db.Model(&models.Task{}).
		Where("user_id = ?", userID).
		Find(&result).Error; err != nil {
		return result, err
	}

	return result, nil
}

func (t *TaskRepository) GetNumberOfUserTaskOnToday(userID int) (int, error) {
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
