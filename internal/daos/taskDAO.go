package daos

import (
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/database"
	models "github.com/manabie-com/togo/internal/models"
)

type TaskDAO struct {
	Host      string
	Port      int
	User      string
	Password  string
	Database  string
	TableName string
}

func (u *TaskDAO) CreateTask(task models.Task) (*models.Task, error) {
	db, err := database.ConnectToDB()
	if err != nil {
		return nil, err
	}
	err = db.Debug().Model(&models.Task{}).Create(&task).Error
	if err != nil {
		return nil, err
	}
	return &task, err
}

func (u *TaskDAO) CountTaskByAccountIDAndPeriod(accountID uuid.UUID, startDate time.Time, endDate time.Time) (uint, error) {
	db, err := database.ConnectToDB()
	if err != nil {
		return 0, err
	}
	tasks := []models.Task{}
	result := db.Find(&tasks, "account_id=? AND created_at BETWEEN ? AND ?", accountID, startDate, endDate)
	return uint(result.RowsAffected), nil
}

func (u *TaskDAO) GetTasksByPeriod(startDate time.Time, endDate time.Time) ([]models.Task, error) {
	db, err := database.ConnectToDB()
	if err != nil {
		return nil, err
	}
	tasks := []models.Task{}
	_ = db.Debug().Find(&tasks, "created_at BETWEEN ? AND ?", startDate, endDate).Error
	return tasks, nil
}
