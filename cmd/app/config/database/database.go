package database

import (
	"fmt"
	"github.com/xrexonx/togo/cmd/app/config/environment"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"time"
)

const (
	_dbConnectionSuccess = "Database connection was successful!"
	_dbConnectionFail    = "Could not connect to database"
	_dbDNS               = "%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

// Config to maintain DB configuration properties
type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

// Instance db instance
var Instance *gorm.DB

// Init initialise database connection from DB configs
func Init() {
	dbEnv := environment.GetDBEnv()
	dbConfig := Config{
		ServerName: dbEnv.DBHost + ":" + dbEnv.DBPort,
		User:       dbEnv.DBUser,
		Password:   dbEnv.DBPass,
		DB:         dbEnv.DBName,
	}
	log.Println("DBConfig:", dbConfig)
	dbConn, err := connect(dbConfig)
	if err != nil {
		log.Fatal(_dbConnectionFail)
	}

	// Connection pooling
	sqlDB, _ := dbConn.DB()
	sqlDB.SetConnMaxLifetime(time.Minute * 30)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)
	//defer sqlDB.Close()

	Instance = dbConn
}

// connect connection to the database
func connect(dbConfig Config) (*gorm.DB, error) {
	var err error
	dsn := fmt.Sprintf(_dbDNS, dbConfig.User, dbConfig.Password, dbConfig.ServerName, dbConfig.DB)
	DBConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, fmt.Errorf(err.Error())
	}
	log.Println(_dbConnectionSuccess)
	return DBConn, nil
}
