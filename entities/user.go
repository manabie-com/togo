package entities

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Plan     string `json:"plan"`
	MaxTodo  int64  `json:"max_todo"`
}

func NewUser() *User {
	return &User{
		Username: "",
		Password: "",
		Plan:     "free",
		MaxTodo:  10,
	}
}

func (u User) IsValid() error {
	if len(u.Username) == 0 {
		return fmt.Errorf("User name is required")
	}
	if len(u.Password) == 0 {
		return fmt.Errorf("User password hash is required")
	}
	return nil
}

func (u *User) PreparePassword() {
	u.Password = hashPassword(u.Password)
}

func hashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func (u User) ComparePassWord(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
