package impl

import (
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/model"
)

type UserRepositoryImpl struct{
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{db}
}

func (r *UserRepositoryImpl) GetUserByIdAndOPassword(id, password string) (model.Users, error) {
	var users model.Users
	if err := r.db.Where("id = ? AND password = ?", id, password).Find(&users).Error; err != nil {
		return model.Users{}, err
	}
	return users, nil
}

func (r *UserRepositoryImpl) GetUserById(id string) (model.Users, error) {
	var users model.Users
	if err := r.db.Where("id = ?", id).Find(&users).Error; err != nil {
		return model.Users{}, err
	}
	return users, nil
}