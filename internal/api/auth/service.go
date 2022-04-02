package auth

import "gorm.io/gorm"

// Crypter represents security interface
type Crypter interface {
	CompareHashAndPassword(hasedPwd string, rawPwd string) bool
	HashPassword(string) string
}

// Auth represents auth application service
type Auth struct {
	db *gorm.DB
	cr Crypter
}

func New(db *gorm.DB, cr Crypter) *Auth {
	return &Auth{
		db: db,
		cr: cr,
	}
}
