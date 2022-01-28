package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var client *gorm.DB

func Initmysql(listmodels ...interface{}) (*gorm.DB, error) {
	if client == nil {
		client, err := gorm.Open(mysql.Open("test.db"), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		client.AutoMigrate(listmodels...)
	}
	return client, nil
}
