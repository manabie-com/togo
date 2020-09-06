package repository

import (
	"time"

	"github.com/google/uuid"
	entity "github.com/manabie-com/togo/internal/entities"

	postgre "github.com/manabie-com/togo/internal/storages/postgre"
)

var (
	TaskRepo taskRepository = &TaskRepository{}
)

type taskRepository interface {
	Add(entity *entity.Task) (*entity.Task, error)
	GetByUserID(userID string, createdDate string) ([]entity.Task, error)
	GetAll(createdDate string) ([]entity.Task, error)
	GetByID(taskID string) (*entity.Task, error)
}

// TaskRepository action CRUD with Task entity
type TaskRepository struct {
	Store *postgre.Storage
}

// Add func add new task
func (repo *TaskRepository) Add(entity *entity.Task) (*entity.Task, error) {
	now := time.Now()

	entity.CreatedDate = now.Format("2006-01-02")

	entity.ID = uuid.New().String()

	result, err := repo.Store.Add(entity)

	return result, err
}

// GetAll func retrives all task
func (repo *TaskRepository) GetAll(createdDate string) ([]entity.Task, error) {

	result, err := repo.Store.GetAll(createdDate)

	return result, err
}

// GetByID func retrives task by primary key
func (repo *TaskRepository) GetByID(taskID string) (*entity.Task, error) {

	result, err := repo.Store.GetByID(taskID)

	return result, err
}

// GetByUserID func retrives task by userID and createdDate
func (repo *TaskRepository) GetByUserID(userID string, createdDate string) ([]entity.Task, error) {

	result, err := repo.Store.GetByUserID(userID, createdDate)

	return result, err
}
