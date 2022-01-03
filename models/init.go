package models

import (
	"gorm.io/gorm"
	"time"
)

type UserConfigs struct {
	gorm.Model
	Email     	string		`gorm:"uniqueIndex:user_day_cfg""`
	Limit 		int8
	Current     int8
	Date		time.Time 	`gorm:"uniqueIndex:user_day_cfg""`
}

func (UserConfigs) TableName() string {
	return "user_configs"
}

type Tasks struct {
	gorm.Model
	Email 	string	`gorm:`
	Task 	string
}

func (Tasks) TableName() string {
	return "tasks"
}
