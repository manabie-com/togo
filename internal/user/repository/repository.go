package repository

import (
	"context"

	"github.com/manabie-com/togo/pkg/errorx"

	"github.com/manabie-com/togo/model"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUser(context.Context, *model.User) (*model.User, error)
	SaveUser(*gorm.DB, *model.User) error
	UpdateUser(*gorm.DB, *model.User) error
}

type userRepository struct {
	*gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) GetUser(ctx context.Context, u *model.User) (*model.User, error) {
	user := &model.User{}
	if err := r.Where(u).
		First(user).
		Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errorx.ErrUserNotFound(err)
		}
		return nil, errorx.ErrDatabase(err)
	}
	return user, nil
}

func (r *userRepository) SaveUser(tx *gorm.DB, user *model.User) error {
	if err := tx.Create(user).Error; err != nil {
		return errorx.ErrDatabase(err)
	}
	return nil
}

func (r *userRepository) UpdateUser(tx *gorm.DB, user *model.User) error {
	if err := tx.Updates(user).Error; err != nil {
		return errorx.ErrDatabase(err)
	}
	return nil
}
