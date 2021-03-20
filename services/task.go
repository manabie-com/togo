package services

import (
	"github.com/manabie-com/togo/models"
	"github.com/manabie-com/togo/repositories"
)

type ITaskService interface {
	GetTaskById(id uint64) (*models.Task, error)
	GetTasksByUserName(username string, createdAt string) (*[]models.Task, error)
	CreateTask(task *models.Task) (*models.Task, error)
	Count(username string) (int64, error)
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

func (taskService *TaskService) GetTasksByUserName(username string, createdAt string) (*[]models.Task, error) {
	return taskService.TaskRepo.GetTasksByUserName(username, createdAt)
}

func (taskService *TaskService) CreateTask(task *models.Task) (*models.Task, error) {
	return taskService.TaskRepo.CreateTask(task)
}

func (taskService *TaskService) Count(username string) (int64, error) {
	return taskService.TaskRepo.Count(username)
}
