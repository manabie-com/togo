package repositories

import (
	"gorm.io/gorm"
)

// LiteDB for working with user
type LiteDB struct {
	DB *gorm.DB
}
