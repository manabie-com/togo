package repositories

import (
	"context"

	"github.com/manabie-com/togo/internal/models"
	"github.com/manabie-com/togo/pkg/txmanager"
	"gorm.io/gorm"
)

//go:generate mockgen -destination=./mock_$GOFILE -source=$GOFILE -package=repositories
type UserRepository interface {
	ValidateUser(ctx context.Context, userID string) (*models.User, error)
}

func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct{ db *gorm.DB }

// ValidateUser returns tasks if match userID AND password
func (r *userRepository) ValidateUser(ctx context.Context, userID string) (*models.User, error) {
	var db = txmanager.GetTxFromContext(ctx, r.db)
	user := models.User{}
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
