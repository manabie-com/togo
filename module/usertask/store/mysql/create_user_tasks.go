package mysql

import (
	"context"
	"togo/module/usertask/model"
)

func (u *userTaskSQL) CreateUserTasks(ctx context.Context, data []model.CreateUserTask) error {
	tx := u.db.Table(model.CreateUserTask{}.TableName()).Begin()
	for _, task := range data {
		task.Status = 1
		if err := tx.Create(&task).Error; err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}