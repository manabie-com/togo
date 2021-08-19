package models

import (
	"gorm.io/gorm"
	"github.com/google/uuid"
)

// Task reflects tasks in DB
type Task struct {
	gorm.Model
	ID          uuid.UUID `json:"id,primary_key"`
	Content     string `json:"content"`
	UserID      uuid.UUID `json:"user_id"`
	CreatedDate string `json:"created_date"`
}