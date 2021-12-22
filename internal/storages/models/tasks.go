package models

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Task struct {
	gorm.Model

	Uid  float64        `gorm:"index" validate:"required"`
	Data datatypes.JSON `validate:"required"`
}

func (_ Task) TableName() string {
	return "tasks"
}
