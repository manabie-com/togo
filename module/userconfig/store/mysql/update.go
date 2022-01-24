package mysql

import (
	"context"
	"togo/module/userconfig/model"
)

func (u *userConfigSQL) Update(ctx context.Context, cond map[string]interface{}, data *model.UpdateUserConfig) error {
	var userCfg model.UserConfig
	db := u.db.Table(userCfg.TableName()).Begin()
	if err := db.Where(cond).Updates(data).Error; err != nil {
		db.Rollback()
		return err
	}

	if err := db.Commit().Error; err != nil {
		return err
	}

	return nil
}