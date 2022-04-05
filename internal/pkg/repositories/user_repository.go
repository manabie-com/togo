package repositories

import (
	"context"
	"togo/internal/pkg/domain/entities"

	"gorm.io/gorm"
)

// UserRepository interface
type UserRepository interface {
	FindUserWithID(ctx context.Context, userID string) (entities.User, error)
	FindUserWithEmail(ctx context.Context, email string) (entities.User, error)
}

type userRepository struct {
	DB *gorm.DB
}

func (r *userRepository) FindUserWithID(ctx context.Context, userID string) (entities.User, error) {
	user := entities.User{}
	result := r.DB.WithContext(ctx).First(&user, userID)
	return user, result.Error
}

func (r *userRepository) FindUserWithEmail(ctx context.Context, email string) (entities.User, error) {
	user := entities.User{}
	result := r.DB.WithContext(ctx).Where("email = ?", email).Take(&user)
	return user, result.Error
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}
