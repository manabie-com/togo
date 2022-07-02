package database

import (
	"fmt"
	"github.com/xrexonx/togo/cmd/app/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

// Config to maintain DB configuration properties
type Config struct {
	ServerName string
	User       string
	Password   string
	DB         string
}

// Init initialise database connection from DB configs
func Init() *gorm.DB {
	dbEnv := config.GetDBEnv()
	dbConfig := Config{
		ServerName: dbEnv.DBHost + ":" + dbEnv.DBPort,
		User:       dbEnv.DBUser,
		Password:   dbEnv.DBPass,
		DB:         dbEnv.DBName,
	}
	log.Println("dbConfig:", dbConfig)
	dbConn, err := connect(dbConfig)
	if err != nil {
		log.Fatal("Could not connect to database")
	}

	// Connection pooling
	//sqlDB, _ := dbConn.DB()
	//sqlDB.SetConnMaxLifetime(time.Minute * 30)
	//sqlDB.SetMaxIdleConns(10)
	//sqlDB.SetMaxOpenConns(10)
	//defer sqlDB.Close()

	// Create tables

	return dbConn
}

// connect connection to the database
func connect(dbConfig Config) (*gorm.DB, error) {
	var err error
	dsn := getConnectionString(dbConfig)
	DBConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	log.Println("Connection was successful!!")
	return DBConn, nil
}

// getConnectionString setup database connection string
func getConnectionString(config Config) string {
	_dns := "%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	connectionString := fmt.Sprintf(_dns, config.User, config.Password, config.ServerName, config.DB)
	return connectionString
}
