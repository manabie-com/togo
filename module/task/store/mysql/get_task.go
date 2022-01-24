package mysql

import (
	"context"
	"togo/module/task/model"
)

func (u *taskSQL) Get(ctx context.Context, cond map[string]interface{}) (*model.Task, error) {
	var user model.Task
	db := u.db.Table(user.TableName())
	if err := db.Where(cond).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}