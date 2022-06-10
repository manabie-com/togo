package mysql

import (
	"context"
	"togo/module/user/model"
)

func (u *userMongo) Get(ctx context.Context, cond map[string]interface{}) (*model.User, error) {
	var user model.User
	db := u.db.Table(user.TableName())
	if err := db.Where(cond).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
