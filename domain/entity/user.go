package entity

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	MaxTodo   int64     `json:"max_todo" db:"max_todo"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

func (u *User) HashPassword() error {
	b, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err == nil {
		u.Password = string(b)
	}
	return err
}

func (u *User) ComparePassword(password string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil, err
}
