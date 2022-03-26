package storage

import (
	"context"

	"github.com/luongdn/togo/models"
	"gorm.io/gorm"
)

func NewUserStore(db *gorm.DB) *userStore {
	return &userStore{
		sqlDB: db,
	}
}

type userStore struct {
	sqlDB *gorm.DB
}

func (s *userStore) CreateUser(ctx context.Context, user *models.User) error {
	result := s.sqlDB.Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
