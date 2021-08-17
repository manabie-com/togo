package storage

import (
	"context"
	"github.com/manabie-com/togo/auth/model"
	"github.com/manabie-com/togo/shared"
	"gorm.io/gorm"
	"strings"
)

func (as *authStorage) FindUserById(ctx context.Context, userId string) (*model.User, error) {
	var user model.User
	db := as.db

	db = db.Where("id = ?", strings.Trim(userId, " "))
	if err := db.First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, shared.ErrNotFound
		}
		return nil, err
	}
	return &user, nil
}
