package repositories

import (
	"context"
	"fmt"
	"togo/database/datastore"
	"togo/domain/models"
)

type UserRepository struct{}

var userRepository *UserRepository

func NewUserRepository() *UserRepository {
	if userRepository == nil {
		userRepository = &UserRepository{}
	}
	return userRepository
}

func (tr *UserRepository) GetByID(ctx context.Context, userID int) (*models.User, error) {
	var user *models.User
	if err := datastore.GetDB().WithContext(ctx).Where("id = ?", userID).First(user).Error; err != nil {
		return nil, fmt.Errorf("error UserRepository.GetByID userID: %d, err: %v", userID, err)
	}
	return user, nil
}
