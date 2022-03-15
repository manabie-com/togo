package models

import (
	"gorm.io/gorm"
)

// Setting for user
type Setting struct {
	gorm.Model
	QuotaPerDay uint64 `json:"quota_per_day"`
	UserID      uint64 `json:"user_id"`
}
