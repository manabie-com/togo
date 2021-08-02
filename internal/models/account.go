package models

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Account model struct
type Account struct {
	ID                 uuid.UUID `gorm:"size:255;column:id;not null;unique; primaryKey;" json:"id"`
	Username           string    `gorm:"size:55;unique; not null" json:"username"`
	Password           string    `gorm:"size:255;not null" json:"password,omitempty"`
	MaxDailyTasksCount uint      `gorm:"column:max_daily_tasks_count;default:5" json:"max_daily_tasks_count,omitempty"`
	Tasks              []Task    `gorm:"foreignKey:AccountID" json:"tasks"`
	//default timestamps
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime" json:"updated_at"`
}

//Hash password
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerifyPassword .
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

//BeforeSave checks Hash
func (u *Account) BeforeSave(*gorm.DB) error {
	if len(u.Password) != 0 {
		hashedPassword, err := Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}

func (u *Account) BeforeUpdate(*gorm.DB) error {
	if len(u.Password) != 0 {
		hashedPassword, err := Hash(u.Password)
		if err != nil {
			return err
		}
		u.Password = string(hashedPassword)
	}
	return nil
}
