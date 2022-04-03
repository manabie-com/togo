package auth

import (
	"gorm.io/gorm"
)

// Crypter represents security interface
type Crypter interface {
	CompareHashAndPassword(hasedPwd string, rawPwd string) bool
	HashPassword(string) string
}

// JWT represents JWT generator interface
type JWT interface {
	GenerateToken(map[string]interface{}) (string, int, error)
}

// Auth represents auth application service
type Auth struct {
	db  *gorm.DB
	cr  Crypter
	jwt JWT
}

// New creates new auth service
func New(db *gorm.DB, cr Crypter, jwt JWT) *Auth {
	return &Auth{
		db:  db,
		cr:  cr,
		jwt: jwt,
	}
}
