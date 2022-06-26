package task

import (
	"example.com/m/v2/internal/models"
	"example.com/m/v2/internal/pkg/id"
	"example.com/m/v2/internal/usecases/task"

	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type repository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) task.TaskUseCase {
	return &repository{
		DB: db,
	}
}

func NewTask(task models.Task) *models.Task {
	return &models.Task{
		ID:         id.NewID().String(),
		Content:    task.Content,
		CreateDate: task.CreateDate,
		UserID:     task.UserID,
	}
}

// AddTask implements task.TaskUseCase
func (r *repository) AddTask(task *models.Task) error {
	if task == nil {
		return errors.New("Invalid input")
	}

	if err := r.DB.Create(task).Error; err != nil {
		return errors.Wrap(err, "Fail Create Task")
	}

	return nil
}

// FindTaskByUser implements task.TaskUseCase
func (r *repository) FindTaskByUser(userID string, createDate string) ([]models.Task, error) {
	if userID == "" || createDate == "" {
		return nil, errors.New("Invalid input")
	}

	tasks := []models.Task{}
	if err := r.DB.Where("user_id = ? and create_date = ?", userID, createDate).Find(&tasks).Error; err != nil {
		return nil, errors.Wrap(err, "Fail query task")
	}

	if len(tasks) == 0 {
		return nil, nil
	}

	return tasks, nil
}