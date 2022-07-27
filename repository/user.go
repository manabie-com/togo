package repository

import (
	"togo/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}

func (u *UserRepository) WithTrx(trxHandle *gorm.DB) *UserRepository {
	if trxHandle == nil {
		return u
	}
	u.db = trxHandle
	return u
}

func (u *UserRepository) Create(user *models.User) (*models.User, error) {
	if err := u.db.Create(user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetByID(userID int) (*models.User, error) {
	var user models.User
	if err := u.db.First(&user, userID).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepository) GetByName(name string) (*models.User, error) {
	var user models.User
	if err := u.db.
		Where("name = ?", name).
		First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
