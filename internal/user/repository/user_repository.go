package repository

import (
	"gorm.io/gorm"
	"togo/domain"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (repo *UserRepository) Create(user domain.User) (domain.User, error) {
	result := repo.db.Create(&user)

	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	result := repo.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return user, nil
}

func (repo *UserRepository) FindById(id int) (domain.User, error) {
	var user domain.User
	result := repo.db.First(&user, id)
	if result.Error != nil {
		return domain.User{}, result.Error
	}
	return user, nil
}
