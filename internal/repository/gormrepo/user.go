package gormrepo

import (
	"context"
	"togo/internal/domain"
	"togo/internal/repository"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository repository constructor
func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepository{
		db,
	}
}

// Create method to create a user
func (r userRepository) Create(ctx context.Context, entity *domain.User) (*domain.User, error) {
	if err := r.db.Create(entity).Error; err != nil {
		return nil, err
	}
	return entity, nil
}

// GetOne method to get one user
func (r userRepository) FindOne(ctx context.Context, filter *domain.User) (*domain.User, error) {
	user := new(domain.User)
	if err := r.db.Where(filter).First(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateByID method to update a user by ID
func (r userRepository) UpdateByID(ctx context.Context, id uint, update *domain.User) (*domain.User, error) {
	user := new(domain.User)
	if err := r.db.Model(user).Updates(update).Error; err != nil {
		return nil, err
	}
	return user, nil
}
