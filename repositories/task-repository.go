package repositories

import (
	"time"

	"github.com/manabie-com/togo/entities"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetAllTask() ([]entities.Task, error)
	CreateTask(task *entities.Task) error
	CountTaskByUserIdAndCreatedAt(userId uint64, createdAt time.Time) (int64, error)
}

type taskConnection struct {
	connection *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &taskConnection{
		connection: db,
	}
}

func (taskConn *taskConnection) GetAllTask() ([]entities.Task, error) {
	var tasks []entities.Task
	err := taskConn.connection.Find(&tasks).Error

	if err != nil {
		return nil, err
	}
	return tasks, nil
}

func (taskConn *taskConnection) CreateTask(task *entities.Task) error {
	err := taskConn.connection.Create(task).Error

	if err != nil {
		return err
	}

	return nil
}

func (taskConn *taskConnection) CountTaskByUserIdAndCreatedAt(userId uint64, createdAt time.Time) (int64, error) {
	var tasks []entities.Task
	var count int64 = 0
	err := taskConn.connection.
		Where("user_id = ? AND date(created_at) = date(?)", userId, createdAt).
		Find(&tasks).Count(&count).
		Error

	if err != nil {
		return 0, err
	}

	return count, nil
}
