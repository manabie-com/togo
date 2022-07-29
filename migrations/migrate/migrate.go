package migrate

import (
	"togo/internal/models"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func Migrate(DB *gorm.DB) error {
	migrate := gormigrate.New(DB, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "1",
			Migrate: func(tx *gorm.DB) error {
				err := tx.AutoMigrate(
					&models.User{},
					&models.Task{},
				)
				return err
			},
			Rollback: func(tx *gorm.DB) error {
				err := tx.Migrator().DropTable(
					&models.User{},
					&models.Task{},
				)
				return err
			},
		},
	})

	err := migrate.Migrate()
	return err
}
