package entities

import (
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
	"time"
)

type CommonModel struct {
	ID        uint           `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

type Task struct {
	CommonModel
	Content string `json:"content" validate:"required"`
	UserID  uint   `json:"user_id" gorm:"not null" validate:"required"`
	Date    string `json:"date" validate:"required" gorm:"not null"`
}

func NewValidator() *validator.Validate {

	v := validator.New()

	return v
}

func (c *Task) Validate() error {

	v := NewValidator()

	err := v.Struct(c)

	return err
}
