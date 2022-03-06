package models

import (
	"database/sql"
	"time"

	"togo/pkg/e"
)

type Task struct {
	ID        int64         `gorm:"column:id;primary_key;auto_increment" db:"id"`
	Name      string        `gorm:"column:name" db:"name"`
	UserID    int           `gorm:"column:member_id" db:"member_id"`
	Status    int           `gorm:"column:status" db:"status"`
	CreatedAt time.Time     `gorm:"column:created_at" db:"created_at"`
	CreatedBy sql.NullInt64 `gorm:"column:created_by;default:null" db:"created_by"`
	UpdatedAt sql.NullTime  `gorm:"column:updated_at;default:null" db:"updated_at"`
	UpdatedBy sql.NullInt64 `gorm:"column:updated_by;default:null" db:"updated_by"`
	DeletedAt sql.NullTime  `gorm:"column:deleted_at;default:null" db:"deleted_at"`
	DeletedBy sql.NullInt64 `gorm:"column:deleted_by;default:null" db:"deleted_by"`
}

func (Task) TableName() string {
	return e.TaskTable
}
