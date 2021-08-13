package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"os"
)

//var db *gorm.DB

func GetPostgersDB() (*gorm.DB, error) {
	//port := `mapstructure:"DB_DRIVER"`
	//user := os.Getenv("APP_DB_USERNAME")
	//password := os.Getenv("APP_DB_PASSWORD")
	//databaseName := os.Getenv("APP_DB_NAME")

	port := "2345"
	user := "postgres"
	password := "Admin@123#"
	databaseName := "postgres"

	desc := fmt.Sprintf("host=localhost port=%s user=%s password=%s dbname=%s sslmode=disable", port, user, password, databaseName)

	db, err := createConnection(desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}

func createConnection(desc string) (*gorm.DB, error) {
	var err error
	var db *gorm.DB
	db, err = gorm.Open("postgres", desc)

	if err != nil {
		return nil, err
	}

	return db, nil
}
