package services

import "github.com/manabie-com/togo/internal/repositories"

type TaskService interface {
}

type taskService struct {
	repo *repositories.Repository
}

func newTaskService(repo *repositories.Repository) TaskService {
	return &taskService{
		repo: repo,
	}
}
