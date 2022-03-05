package migration

import (
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"

	// Register using Golang migrate.
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"gorm.io/gorm"

	"github.com/khangjig/togo/client/logger"
	"github.com/khangjig/togo/config"
)

func Up(db *gorm.DB) {
	getDB, err := db.DB()
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}

	driver, err := mysql.WithInstance(getDB, &mysql.Config{MigrationsTable: "migration"})
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}

	m, err := migrate.NewWithDatabaseInstance("file://./migration", config.GetConfig().MySQL.DBName, driver)
	if err != nil {
		logger.GetLogger().Fatal(err.Error())
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		logger.GetLogger().Fatal(err.Error())
	}

	logger.GetLogger().Info("Up done!")
}
