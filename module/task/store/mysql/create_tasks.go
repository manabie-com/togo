package mysql

import (
	"context"
	"togo/module/task/model"
)

func (u *taskSQL) CreateTasks(ctx context.Context, data []model.CreateTask) error {
	tx := u.db.Table(model.Task{}.TableName()).Begin()
	for i, _ := range data {
		data[i].Status = 1
		err := tx.Create(&data[i]).Error
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}