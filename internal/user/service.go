package user

import (
	"github.com/xrexonx/togo/internal/repository"
	"gorm.io/gorm"
)

func GetById(db *gorm.DB, id string) User {
	var user, _ = repository.GetByID[User](db, id)
	return user
}
