package repositories

import (
	"gorm.io/gorm"
)

type UserRepository interface {
}

func newUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

type userRepository struct{ db *gorm.DB }
