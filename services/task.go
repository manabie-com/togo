package services

import (
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/repositories"
)

type ITaskService interface {
	GetTaskById(id uint64) (*models.Task, error)
	GetTasks() (*[]models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
}

type TaskService struct {
	TaskRepo repositories.ITaskRepository
}

func NewTaskService(taskRepository *repositories.ITaskRepository) ITaskService {
	return &TaskService{TaskRepo: *taskRepository}
}

func (taskService *TaskService) GetTaskById(id uint64) (*models.Task, error) {
	return taskService.TaskRepo.GetTaskById(id)
}

func (taskService *TaskService) GetTasks() (*[]models.Task, error) {
	return taskService.TaskRepo.GetTasks()
}

func (taskService *TaskService) CreateTask(task *models.Task) (*models.Task, error) {
	return taskService.TaskRepo.CreateTask(task)
}
