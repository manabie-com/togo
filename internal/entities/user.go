package entities

import (
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       int    `json:"id"`
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

func (u User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
