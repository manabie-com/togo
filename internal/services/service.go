package services

import "github.com/manabie-com/togo/internal/repositories"

type Service struct {
	TaskService TaskService
	UserService UserService
}

// InitServiceFactory initialize services factory
func InitServiceFactory(repo *repositories.Repository) *Service {

	return &Service{
		TaskService: newTaskService(repo),
		UserService: newUserService(repo),
	}
}
