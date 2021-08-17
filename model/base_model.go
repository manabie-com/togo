package model

import "time"

type BaseModel struct {
	Id        int        `gorm:"column:id;primaryKey"`
	Status    int        `gorm:"column:status;default:1;"`
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
}
