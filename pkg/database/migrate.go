package database

import (
	"errors"
	"togo-service/app/models"
	"togo-service/pkg/utils"

	"gorm.io/gorm"
)

func DoMigrate(db *gorm.DB) {
	// db.AutoMigrate(&models.User{})

	if err := db.AutoMigrate(&models.User{}); err == nil && db.Migrator().HasTable(&models.User{}) {
		if err := db.First(&models.User{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
			//Insert seed data
			pass, _ := utils.HashPassword("secret")
			admin := models.User{
				Username: "admin@admin.com",
				Password: pass,
				Role:     "admin",
			}
			db.Create(&admin)
		}
	}

	db.AutoMigrate(&models.Setting{})
	db.AutoMigrate(&models.Task{})
}
