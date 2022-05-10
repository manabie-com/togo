package entity

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
)

type Task struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	UserID      uint64    `gorm:"size:100;not null;" json:"user_id"`
	Title       string    `gorm:"size:100;not null;" json:"title"`
	Description string    `gorm:"text;not null;" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

func (f *Task) BeforeSave(tx *gorm.DB) error {
	f.Title = html.EscapeString(strings.TrimSpace(f.Title))
	return nil
}

func (f *Task) Prepare() {
	f.Title = html.EscapeString(strings.TrimSpace(f.Title))
	f.CreatedAt = time.Now()
}

func (f *Task) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)

	switch strings.ToLower(action) {
	case "update":
		if f.Title == "" || f.Title == "null" {
			errorMessages["title_required"] = "title is required"
		}
		if f.Description == "" || f.Description == "null" {
			errorMessages["desc_required"] = "description is required"
		}
	default:
		if f.Title == "" || f.Title == "null" {
			errorMessages["title_required"] = "title is required"
		}
		if f.Description == "" || f.Description == "null" {
			errorMessages["desc_required"] = "description is required"
		}
	}
	return errorMessages
}
