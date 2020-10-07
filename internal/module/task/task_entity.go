package task

import "time"

// define role
const (
	StatusActive string = "active"
	StatusDelete string = "delete"
)

// Task entity
type Task struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Content     string    `gorm:"type:varchar(2048)" json:"content"`
	Status      string    `gorm:"type:varchar(64)" json:"status"`
	UserID      uint64    `gorm:"type:integer" json:"user_id"`
	CreatedDate time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_date"`
}
