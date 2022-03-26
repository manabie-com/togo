package models

import (
	"database/sql/driver"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// Customize data type for `Action` field
type UserAction string

const (
	TaskCreate UserAction = "tasks/create"
)

func (action *UserAction) Scan(value interface{}) error {
	*action = UserAction(value.([]byte))
	return nil
}

func (action UserAction) Value() (driver.Value, error) {
	return string(action), nil
}

// Customize data type for `Unit` field
type TimeUnit string

const (
	Second TimeUnit = "second"
	Minute TimeUnit = "minute"
	Hour   TimeUnit = "hour"
	Day    TimeUnit = "day"
)

func (unit *TimeUnit) Scan(value interface{}) error {
	*unit = TimeUnit(value.([]byte))
	return nil
}

func (unit TimeUnit) Value() (driver.Value, error) {
	return string(unit), nil
}

type Rule struct {
	gorm.Model
	ID              string `json:"id" gorm:"size:191"`
	UserID          string `json:"user_id" gorm:"primary_key;size:191"`
	User            User
	Action          UserAction `json:"action" gorm:"primary_key" sql:"Action"`
	Unit            TimeUnit   `json:"unit" gorm:"type:text" sql:"TimeUnit"`
	RequestsPerUnit int64      `json:"requests_per_unit"`
}

func (rule *Rule) BeforeCreate(tx *gorm.DB) (err error) {
	if _, err := uuid.Parse(rule.ID); err != nil {
		rule.ID = uuid.NewString()
	}
	return nil
}
