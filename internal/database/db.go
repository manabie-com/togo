package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"strings"
)

var db *gorm.DB
var dbSettings DBSettings

// DBSettings !
type DBSettings struct {
	Type     string `json:"type"` // sqlite, mysql, postgres
	Name     string `json:"name"` // File/DB name
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

func InitializeDB(settings DBSettings, models ...interface{}) {
	// set settings to dbSettings
	dbSettings = settings
	// Open the connection the the DB
	db = GetDB()

	// Migrate schema
	// db.AutoMigrate(models)
	for _, model := range models {
		db.AutoMigrate(model)
	}
}

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	var err error

	if strings.ToLower(dbSettings.Type) == "sqlite" {
		dbName := dbSettings.Name
		if dbName == "" {
			dbName = "test.db"
		}
		db, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
		if err != nil {
			fmt.Errorf(err.Error())
		}
	} else if strings.ToLower(dbSettings.Type) == "mysql" {
		if dbSettings.Host == "" || dbSettings.Host == "localhost" {
			dbSettings.Host = "127.0.0.1"
		}
		if dbSettings.Port == 0 {
			dbSettings.Port = 3306
		}

		if dbSettings.User == "" {
			dbSettings.User = "root"
		}

		credential := dbSettings.User

		if dbSettings.Password != "" {
			credential = fmt.Sprintf("%s:%s", dbSettings.User, dbSettings.Password)
		}
		dsn := fmt.Sprintf("%s@(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
			credential,
			dbSettings.Host,
			dbSettings.Port,
			dbSettings.Name,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			fmt.Errorf(err.Error())
		}

		// Check if the error is DB doesn't exist and create it
		if err != nil && err.Error() == "Error 1049: Unknown database '"+dbSettings.Name+"'" {
			err = createMysqlDB()

			if err == nil {
				db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
				if err != nil {
					fmt.Errorf(err.Error())
				}
			}
		}
	}
	return db
}

func createMysqlDB() error {
	credential := dbSettings.User

	if dbSettings.Password != "" {
		credential = fmt.Sprintf("%s:%s", dbSettings.User, dbSettings.Password)
	}

	dsn := fmt.Sprintf("%s@(%s:%d)/?charset=utf8&parseTime=True&loc=Local",
		credential,
		dbSettings.Host,
		dbSettings.Port,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	// create database
	db = db.Exec("CREATE SCHEMA `" + dbSettings.Name + "` DEFAULT CHARACTER SET utf8 COLLATE utf8_bin")

	if db.Error != nil {
		return fmt.Errorf(db.Error.Error())
	}

	return nil
}
