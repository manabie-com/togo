package users

import (
	"time"
	"togo/src/common/bcrypt"
	"togo/src/common/types"

	"gorm.io/gorm"
)

type User struct {
	ID        uint      `json:"id" gorm:"col:id,primary_key"`
	Username  string    `json:"username" gorm:"col:username"`
	Password  string    `json:"password" gorm:"col:password"`
	CreatedAt time.Time `json:"created_at" gorm:"col:created_at"`
	UpdatedAt time.Time `json:"updated_at" gorm:"col:updated_at"`
}

func (u *User) ToJSON() types.JSON {
	return types.JSON{
		"id":         u.ID,
		"username":   u.Username,
		"created_at": u.CreatedAt,
		"updated_at": u.UpdatedAt,
	}
}

func (u *User) Read(m types.JSON) {
	u.ID = uint(m["id"].(float64))
	u.Username = m["username"].(string)
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	var hash, hashErr = bcrypt.HashPassword(u.Password)

	if hashErr != nil {
		return hashErr
	}

	u.Password = hash
	return
}
