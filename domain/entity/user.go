package entity

import (
	"html"
	"strings"
	"time"

	"github.com/jfzam/togo/infrastructure/security"
)

type User struct {
	ID              uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserName        string    `gorm:"size:50;not null;" json:"username"`
	Password        string    `gorm:"size:50;not null;" json:"password"`
	TaskLimitPerDay int64     `gorm:"default:0" json:"task_limit_per_day"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type PublicUser struct {
	ID       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserName string `gorm:"size:50;not null;" json:"username"`
}

// BeforeSave hash the password
func (u *User) BeforeSave() error {
	hashPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

type Users []User

// PublicUsers return list of users
func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}

// PublicUser return user
func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:       u.ID,
		UserName: u.UserName,
	}
}

// Prepare function prepares the user data
func (u *User) Prepare() {
	u.UserName = html.EscapeString(strings.TrimSpace(u.UserName))
	u.CreatedAt = time.Now()
}

// Validate validates user action
func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	//var err error

	switch strings.ToLower(action) {
	case "login":
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.UserName == "" {
			errorMessages["email_required"] = "username is required"
		}
		if u.UserName != "" {
			// if err = checkmail.ValidateFormat(u.Email); err != nil {
			// 	errorMessages["invalid_email"] = "please provide a valid email"
			// }
		}
	default:
		if u.UserName == "" {
			errorMessages["firstname_required"] = "first name is required"
		}
		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}
		if u.Password != "" && len(u.Password) < 12 {
			errorMessages["invalid_password"] = "password should be at least 12 characters"
		}
		if u.UserName != "" {
			// if err = checkmail.ValidateFormat(u.Email); err != nil {
			// 	errorMessages["invalid_email"] = "please provide a valid email"
			//
		}
	}
	return errorMessages
}
