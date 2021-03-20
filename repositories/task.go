package repositories

import (
	"github.com/manabie-com/togo/models"
	"gorm.io/gorm"
	"time"
)

type ITaskRepository interface {
	GetTaskById(id uint64) (*models.Task, error)
	GetTasksByUserName(username string, createdAt string) (*[]models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
	Count(username string) (int64, error)
}

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) ITaskRepository {
	return &TaskRepository{DB: db}
}

func (taskRepo *TaskRepository) GetTaskById(id uint64) (*models.Task, error) {
	var task models.Task
	result := taskRepo.DB.First(&task, id)
	return &task, result.Error
}

func (taskRepo *TaskRepository) GetTasksByUserName(username string, createdAt string) (*[]models.Task, error) {
	var tasks []models.Task

	if createdAt != "" {
		result := taskRepo.DB.Where("username = ? ", username).
			Where("DATE(created_at) = DATE(?)", createdAt).Find(&tasks)

		return &tasks, result.Error
	}

	result := taskRepo.DB.Where("username = ? ", username).Find(&tasks)

	return &tasks, result.Error
}

func (taskRepo *TaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	result := taskRepo.DB.Create(task)
	return task, result.Error
}

func (taskRepo *TaskRepository) Count(username string) (int64, error) {
	var count int64

	result := taskRepo.DB.Model(&models.Task{}).
		Where("DATE(created_at) = DATE(?)", time.Now()).
		Where("username = ?", username).
		Count(&count)

	return count, result.Error
}
