package models

import (
	"errors"
	"html"
	"strings"
)

type User struct {
	ID       uint64
	Username string
	Password string
}

// // BeforeSave hash the user password
// func (u *User) BeforeSave() error {
// 	hashedPassword, err := utils.Hash(u.Password)
// 	if err != nil {
// 		return err
// 	}
// 	u.Password = string(hashedPassword)
// 	return nil
// }

// Prepare cleans the inputs
func (u *User) Prepare() {
	u.Password = html.EscapeString(strings.TrimSpace(u.Password))
}

// Validate validates the inputs
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "login":
		if u.Username == "" {
			return errors.New("username is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
	default:
		if u.Username == "" {
			return errors.New("username is required")
		}
		if u.Password == "" {
			return errors.New("password is required")
		}
	}
	return nil
}
