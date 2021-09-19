package repositories

import (
	"gorm.io/gorm"
)

type Repository struct {
	TaskRepository TaskRepository
	UserRepository UserRepository
}

// InitRepositoryFactory init repositories factory
func InitRepositoryFactory(db *gorm.DB) *Repository {
	return &Repository{
		TaskRepository: newTaskRepository(db),
		UserRepository: newUserRepository(db),
	}
}
