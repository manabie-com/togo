package service

import (
	"time"
	"togo/domain"
)

type TaskService struct {
	repo     domain.ITaskRepository
	userRepo domain.IUserRepository
}

func NewTaskService(repo domain.ITaskRepository, userRepo domain.IUserRepository) domain.ITaskService {
	return &TaskService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (service *TaskService) Create(params domain.TaskParams) (domain.Task, error) {
	task := domain.Task{
		Content:   params.Content,
		CreatedAt: time.Time{},
	}

	newTask, err := service.repo.Create(task)
	if err != nil {
		return domain.Task{}, err
	}

	return newTask, nil
}
