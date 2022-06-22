package entities

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	PasswordHash string `json:"password_hash"`
	Plan         string `json:"plan"`
	MaxTodo      int64  `json:"max_todo"`
}

func (u User) IsValid() error {
	if len(u.Name) == 0 {
		return fmt.Errorf("User name is required")
	}
	if len(u.PasswordHash) == 0 {
		return fmt.Errorf("User password hash is required")
	}
	return nil
}

func (u *User) PreparePassword() {
	u.PasswordHash = hashPassword(u.PasswordHash)
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func (u User) ComparePassWord(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.PasswordHash), []byte(password))
	return err == nil
}
