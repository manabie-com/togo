package tasks

import (
	"context"
	"database/sql"

	"log"

	"github.com/manabie-com/togo/internal/models"
	timeUtils "github.com/manabie-com/togo/internal/utils/time"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

// Create new task service
func NewService(db *sql.DB) ITaskService {
	return &TaskService{
		DB: db,
	}
}

// Task service interface
type ITaskService interface {
	CreateNew(context.Context, *models.Task) error
	GetTasksCreatedOn(context.Context, string) []*models.Task
	TaskCountByUser(context.Context, string) (int64, error)
}

// Task service object
type TaskService struct {
	DB *sql.DB
}

// Create new task
func (t TaskService) CreateNew(ctx context.Context, task *models.Task) error {
	return task.Insert(ctx, t.DB, boil.Infer())
}

// Count number of tasks
func (t TaskService) TaskCountByUser(ctx context.Context, userId string) (int64, error) {
	return models.Tasks(
		models.TaskWhere.UserID.EQ(userId),
		models.TaskWhere.CreateDate.EQ(timeUtils.CurrentDate()),
	).Count(ctx, t.DB)

}

// Retrieve task
func (t TaskService) GetTasksCreatedOn(ctx context.Context, date string) []*models.Task {
	data, err := models.Tasks().All(ctx, t.DB)
	if err != nil {
		log.Println(err)
		return []*models.Task{}
	}
	return data
}
