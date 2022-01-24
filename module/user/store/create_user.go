package store

import (
	"context"
	"togo/module/user/model"
)

func (u *userSQL) Create(ctx context.Context, data *model.CreateUser) error {
	data.Status = 1

	tx := u.db.Table(data.TableName()).Begin()
	if err := tx.Create(&data).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}