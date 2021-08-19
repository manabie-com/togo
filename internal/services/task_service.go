package services

import (
	models "github.com/manabie-com/togo/internal/models"
	repositories "github.com/manabie-com/togo/internal/repositories"
)

type TaskService struct {
	TaskRepository repositories.TaskRepository
}

func ProvideTaskService(repo repositories.TaskRepository) TaskService {
	return TaskService{TaskRepository: repo}
}

func (p *TaskService) FindAll() []models.Task {
	return p.TaskRepository.FindAll()
}

func (p *TaskService) FindByID(id string) models.Task {
	return p.TaskRepository.FindByID(id)
}

func (p *TaskService) Save(task models.Task) models.Task {
	p.TaskRepository.Save(task)

	return task
}

func (p *TaskService) Delete(task models.Task) {
	p.TaskRepository.Delete(task)
}
