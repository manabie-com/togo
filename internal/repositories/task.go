package repositories

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/pkg/txmanager"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=./mock_$GOFILE -source=$GOFILE -package=repositories
type TaskRepository interface {
	ListTasks(ctx context.Context, userID string, createDate time.Time) ([]models.Task, error)
	AddTask(ctx context.Context, task models.Task) (string, error)
	GetTask(ctx context.Context, id string) (*models.Task, error)
}

func newTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db}
}

type taskRepository struct{ db *gorm.DB }

// ListTasks returns tasks if match userID AND createDate.
func (r *taskRepository) ListTasks(ctx context.Context, userID string, createDate time.Time) ([]models.Task, error) {
	var db = txmanager.GetTxFromContext(ctx, r.db)
	var tasks []models.Task
	startTime := time.Date(createDate.Year(), createDate.Month(), createDate.Day(), 0, 0, 0, 0, time.Local)
	endTime := time.Date(createDate.Year(), createDate.Month(), createDate.Day(), 23, 59, 59, 999, time.Local)
	if err := db.Where("user_id = ? AND created_date >= ? AND created_date <= ?", userID, startTime, endTime).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

// AddTask adds a new task to DB
func (r *taskRepository) AddTask(ctx context.Context, task models.Task) (string, error) {
	var (
		db  = txmanager.GetTxFromContext(ctx, r.db)
		id  string
		now = time.Now()
	)

	task.ID = uuid.New().String()
	task.CreatedDate = now
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
	if err := db.Raw(query, task.ID, task.Content, task.UserID, task.CreatedDate, startTime, endTime).Scan(&id).Error; err != nil {
		return "", err
	}
	return id, nil
}

// GetTask
func (r *taskRepository) GetTask(ctx context.Context, id string) (*models.Task, error) {
	var db = txmanager.GetTxFromContext(ctx, r.db)
	task := models.Task{}
	if err := db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return &task, nil
}
