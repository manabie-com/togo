package mysql

import (
	"context"
	"togo/module/userconfig/model"
)

func (u *userConfigSQL) Get(ctx context.Context, cond map[string]interface{}) (*model.UserConfig, error) {
	var userCfg model.UserConfig
	db := u.db.Table(userCfg.TableName())
	if err := db.Where(cond).First(&userCfg).Error; err != nil {
		return nil, err
	}

	return &userCfg, nil
}