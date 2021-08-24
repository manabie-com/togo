package repositories

import (
	"context"
	"github.com/manabie-com/togo/internal/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

type UserRepository interface {
	ValidateUser(ctx context.Context, userID, password string) bool
}

func NewUserRepository(injectedDB *gorm.DB) UserRepository {
	return &userRepository{
		db: injectedDB,
	}
}

func (r *userRepository) ValidateUser(ctx context.Context, userID, password string) bool {
	var user = &models.User{}
	if err := r.db.Model(user).
		Where("id = ? AND password = ?", userID, password).Error; err != nil {
		return false
	}

	return true
}
