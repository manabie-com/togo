package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	//"os"
)

//var db *gorm.DB

func GetPostgersDB(host, port, user, password, databaseName string) (*gorm.DB, error) {

	desc := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, databaseName)

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
