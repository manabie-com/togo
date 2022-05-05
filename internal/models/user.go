package models

import (
	"github.com/jmsemira/togo/internal/database"
)

type User struct {
	ID              uint
	Username        string
	Password        string
	RateLimitPerDay int
}

func (u *User) UsernameExist() bool {
	db := database.GetDB()
	user := User{}

	db.Where("username = ?", u.Username).First(&user)

	if user.ID > 0 {
		return true
	}
	return false
}

func (u *User) Save() {
	db := database.GetDB()
	db.Create(u)
}

func (u *User) GetUserbyUsername(username string) {
	db := database.GetDB()
	db.Where("username = ?", username).First(u)
}
