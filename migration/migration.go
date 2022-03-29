package migration

import (
	"togo/models"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Task{})
}

func Rollback(db *gorm.DB) {
	db.Migrator().DropTable(&models.User{})
	db.Migrator().DropTable(&models.Task{})
}