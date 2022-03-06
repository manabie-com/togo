package migration

import (
	"togo/pkg/e"
)

type Migration struct {
	ID        uint   `gorm:"column:id;primary_key;auto_increment"`
	Migration string `gorm:"column:migration"`
	Batch     int    `gorm:"column:batch"`
}

func (Migration) TableName() string {
	return e.MigrationTable
}
