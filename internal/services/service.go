package services

import (
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/pkg/txmanager"
)

type Service struct {
	TaskService TaskService
	UserService UserService
}

// InitServiceFactory initialize services factory
func InitServiceFactory(repo *repositories.Repository, tx txmanager.TransactionManager) *Service {
	return &Service{
		TaskService: newTaskService(repo, tx),
		UserService: newUserService(repo),
	}
}
