package services

import (
	"fmt"
	models "github.com/manabie-com/togo/internal/models"
	repositories "github.com/manabie-com/togo/internal/repositories"
)

type TaskService struct {
	TaskRepository repositories.TaskRepository
}

func ProvideTaskService(repo repositories.TaskRepository) TaskService {
	return TaskService{TaskRepository: repo}
}

func (service *TaskService) FindAll() []models.Task {
	result := service.TaskRepository.FindAll()
	fmt.Println("Repo findAll()")
	return result
}

func (service *TaskService) FindByID(id string) models.Task {
	return service.TaskRepository.FindByID(id)
}

func (service *TaskService) Create(task models.Task) models.Task {
	service.TaskRepository.Create(task)

	return task
}

func (service *TaskService) Delete(task models.Task) {
	service.TaskRepository.Delete(task)
}
