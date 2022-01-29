package database

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var client *gorm.DB

func Initmysql(listmodels ...interface{}) (*gorm.DB, error) {
	if client == nil {
		count := 0
		newclient, err := gorm.Open(mysql.Open("user_testdb:password@tcp(test-db:3306)/testdb"), &gorm.Config{})
		if err != nil {
			for {
				if err == nil {
					fmt.Println("")
					break
				}
				fmt.Print(".")
				time.Sleep(time.Second)
				count++
				if count > 180 {
					fmt.Println("")
					fmt.Println("DB connection failure")
					return nil, err
				}
				newclient, err = gorm.Open(mysql.Open("user_testdb:password@tcp(test-db:3306)/testdb"), &gorm.Config{})
			}
		}
		newclient.AutoMigrate(listmodels...)
		client = newclient
	}
	return client, nil
}
