package plan

import "gorm.io/gorm"

// Plan represents plan service
type Plan struct {
	db *gorm.DB
}

// New creates new plan service
func New(db *gorm.DB) *Plan {
	return &Plan{
		db: db,
	}
}
