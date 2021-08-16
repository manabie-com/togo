package storage

import (
	"context"
	"github.com/manabie-com/togo/shared"
	"github.com/manabie-com/togo/user/model"
	"gorm.io/gorm"
)

func (s *userStorage) FindById(_ context.Context, id int) (*model.User, error) {
	var user model.User
	db := s.db

	db = db.Table(user.TableName())
	db = db.Where("id = ?", id)

	if err := db.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, shared.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
