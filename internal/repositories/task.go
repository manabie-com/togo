package repositories

import (
	"context"
	"errors"
	"time"

	"github.com/manabie-com/togo/internal/models"
	"gorm.io/gorm"
)

type TaskRepository interface {
	ListTasks(ctx context.Context, userID string, createDate time.Time) ([]models.Task, error)
	AddTask(ctx context.Context, task models.Task) error
	GetTask(ctx context.Context, id string) (*models.Task, error)
}

func newTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

type taskRepository struct{ db *gorm.DB }

// ListTasks returns tasks if match userID AND createDate.
func (r *taskRepository) ListTasks(ctx context.Context, userID string, createDate time.Time) ([]models.Task, error) {
	var tasks []models.Task
	startTime := time.Date(createDate.Year(), createDate.Month(), createDate.Day(), 0, 0, 0, 0, time.Local)
	endTime := time.Date(createDate.Year(), createDate.Month(), createDate.Day(), 23, 59, 59, 999, time.Local)
	if err := r.db.Where("user_id = ? AND created_date >= ? AND created_date <= ?", userID, startTime, endTime).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// AddTask adds a new task to DB
func (r *taskRepository) AddTask(ctx context.Context, task models.Task) error {
	var id string
	now := time.Now()
	startTime := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.Local)
	endTime := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 999, time.Local)
	query := `
		INSERT INTO tasks (id, content, user_id, created_date)
		SELECT ?, ?, ?, ?
		WHERE (
				SELECT COUNT(*)
				FROM tasks
				WHERE created_date >= ?
					AND created_date <= ?
			) <= 4
		RETURNING id
	`
	if err := r.db.Raw(query, task.ID, task.Content, task.UserID, task.CreatedDate, startTime, endTime).Scan(&id).Error; err != nil {
		return err
	}
	if id == "" {
		return errors.New("limit 5 tasks per day")
	}
	return nil
}

// GetTask
func (r *taskRepository) GetTask(ctx context.Context, id string) (*models.Task, error) {
	task := models.Task{}
	if err := r.db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
