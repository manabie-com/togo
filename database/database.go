package database

import (
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var (
	DB *gorm.DB
)

func InitDB() *gorm.DB {
	DB = ConnectDB()
	return DB
}

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DB_USERNAME") + ":" + os.Getenv("DB_PASSWORD") +
		"@tcp(" + os.Getenv("DB_HOST") + ")/" + os.Getenv("DB_DATABASE") +
		"?charset=utf8&parseTime=True&loc=" + os.Getenv("APP_TIMEZONE")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Error connecting to database", err)
		return nil
	}

	return db
}
