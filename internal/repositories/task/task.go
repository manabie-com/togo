package task

import (
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/internal/pkg/id"
	"github.com/manabie-com/togo/internal/usecases/task"

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
	taskID := id.NewID().String()
	if task.ID != "" {
		taskID = task.ID
	}
	return &models.Task{
		ID:         taskID,
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
