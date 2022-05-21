package domain

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" validate:"required"`
	TaskLimit int       `json:"task_limit"`
	CreatedAt time.Time `json:"created_at"`
}

type UserParams struct {
	Email     string `json:"email"`
	TaskLimit int    `json:"task_limit"`
}

type IUserRepository interface {
	Create(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id int) (User, error)
}
