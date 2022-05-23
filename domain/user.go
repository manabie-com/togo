package domain

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" validate:"required"`
	TaskLimit int       `json:"task_limit"`
	CreatedAt time.Time `json:"created_at"`
}

type IUserRepository interface {
	SetTx(tx *gorm.DB) *gorm.DB
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
	Save(user User) error
}
