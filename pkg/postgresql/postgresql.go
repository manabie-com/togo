package postgresql

import (
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func GetInstance(dsn string) (*gorm.DB, error) {

	gormConfig := &gorm.Config{
		Logger: logger.New(
			logrus.New(), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
		PrepareStmt: true,
	}
	postgresConfig := postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}
	db, err := gorm.Open(postgres.New(postgresConfig), gormConfig)
	if err != nil {
		return nil, err
	}
	// Get generic database object sql.DB to use its functions
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(15)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(5)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Minute * 30)

	return db, nil
}
