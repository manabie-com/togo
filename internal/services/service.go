package services

import (
	"database/sql"

	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/pkg/transactional"
)

type Service struct {
	TaskService TaskService
	UserService UserService
}

// InitServiceFactory initialize services factory
func InitServiceFactory(db *sql.DB, repo *repositories.Repository) *Service {
	_db := transactional.NewDB(db)
	return &Service{
		TaskService: newTaskService(repo, _db),
		UserService: newUserService(repo),
	}
}
