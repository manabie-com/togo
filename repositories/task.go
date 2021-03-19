package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/models"
)

type ITaskRepository interface {
	GetTaskById(id int) (*models.Task, error)
	GetAllTasks() (*[]models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
}

type TaskRepository struct {
	DB *gorm.DB
}

func (taskRepository *TaskRepository) GetTaskById(id uint64) (*models.Task, error) {
	var task models.Task
	result := taskRepository.DB.First(&task, id)
	return &task, result.Error
}

func (taskRepository *TaskRepository) GetAllTasks() (*[]models.Task, error) {
	var tasks []models.Task
	result := taskRepository.DB.Find(&tasks)
	return &tasks, result.Error
}

func (taskRepository *TaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	result := taskRepository.DB.Create(task)
	return task, result.Error
}
