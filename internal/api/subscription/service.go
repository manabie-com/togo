package subscription

import (
	"gorm.io/gorm"
)

// Subscription represents subscription service
type Subscription struct {
	db *gorm.DB
}

// New creates new subscription service
func New(db *gorm.DB) *Subscription {
	return &Subscription{
		db: db,
	}
}
