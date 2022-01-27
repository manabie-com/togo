package repository

import "gorm.io/gorm"

type UserRepository interface {
}

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}
