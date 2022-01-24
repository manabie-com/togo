package mysql

import (
	"context"
	"togo/module/usertask/model"
)

func (u *userTaskSQL) Get(ctx context.Context, cond map[string]interface{}) (*model.UserTask, error) {
	var userTask model.UserTask
	db := u.db.Table(userTask.TableName())
	if err := db.Where(cond).First(&userTask).Error; err != nil {
		return nil, err
	}

	return &userTask, nil
}