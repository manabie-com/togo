package mysql

import (
	"context"
	"togo/module/userconfig/model"
)

func (u *userConfigSQL) Create(ctx context.Context, data *model.CreateUserConfig) error {
	tx := u.db.Table(data.TableName()).Begin()
	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}