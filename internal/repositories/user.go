package repositories

import (
	"context"

	"github.com/manabie-com/togo/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	ValidateUser(ctx context.Context, userID string) (*models.User, error)
}

func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct{ db *gorm.DB }

// ValidateUser returns tasks if match userID AND password
func (r *userRepository) ValidateUser(ctx context.Context, userID string) (*models.User, error) {
	user := models.User{}
	if err := r.db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
