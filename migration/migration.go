package migration

import (
	"github.com/manabie-com/togo/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(
		&models.User{},
		&models.Task{},
	)
	return err
}
