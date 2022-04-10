package db

import (
	"github.com/namnhatdoan/togo/models"
	"github.com/namnhatdoan/togo/settings"
	"gorm.io/gorm"
	gorm_logger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var log = settings.GetLogger()

var db *gorm.DB

func init() {
	var err error
	logLevel := gorm_logger.Silent
	if settings.DbDebug {
		logLevel =gorm_logger.Info
	}
	db, err = gorm.Open(settings.DbDialector, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
		},
		Logger: gorm_logger.Default.LogMode(logLevel),
	})

	if err != nil {
		log.WithField("DB_Name", settings.DbName).WithError(err).Error("Error when init db")
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.WithError(err).Error("Error when init db connection")
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(2)
	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(10)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	migrate()
}

func GetDB() *gorm.DB {
	return db
}

func migrate() {
	db.AutoMigrate(&models.UserConfigs{})
	db.AutoMigrate(&models.Tasks{})
}
