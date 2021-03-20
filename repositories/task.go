package repositories

import (
	"github.com/manabie-com/togo/models"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	GetTaskById(id uint64) (*models.Task, error)
	GetTasks() (*[]models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
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

func (taskRepo *TaskRepository) GetTasks() (*[]models.Task, error) {
	var tasks []models.Task
	result := taskRepo.DB.Find(&tasks)
	return &tasks, result.Error
}

func (taskRepo *TaskRepository) CreateTask(task *models.Task) (*models.Task, error) {
	result := taskRepo.DB.Create(task)
	return task, result.Error
}
