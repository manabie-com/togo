package datastore

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"togo/config"
)

var (
	dbConn *gorm.DB
	err    error
)

// SetupDB opens a database and saves the reference to `Database` struct.
func SetupDB() error {
	var db = dbConn

	configuration := config.GetConfig()

	driver := configuration.Database.Driver
	database := configuration.Database.Dbname
	username := configuration.Database.Username
	password := configuration.Database.Password
	host := configuration.Database.Host
	port := configuration.Database.Port
	charset := configuration.Database.Charset

	var dsn string

	if driver == "mysql" { // MYSQL
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			username,
			password,
			host,
			port,
			database,
			charset,
		)

		newLogger := logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		)

		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			return fmt.Errorf("db connection err: %v", err)
		}
	}

	mySQLConn, err := db.DB()
	if err != nil {
		return fmt.Errorf("extracting mysqlDB from gorm erorr %v", err)
	}

	mySQLConn.SetMaxIdleConns(configuration.Database.MaxIdleConns)
	mySQLConn.SetMaxOpenConns(configuration.Database.MaxOpenConns)
	mySQLConn.SetConnMaxLifetime(time.Duration(configuration.Database.MaxLifetime) * time.Second)
	dbConn = db
	return nil
}

func GetDB() *gorm.DB {
	return dbConn
}

func CloseDB() error {
	db, err := dbConn.DB()
	if err != nil {
		return err
	}

	err = db.Close()
	if err != nil {
		return err
	}
	return nil
}
