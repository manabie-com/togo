package database

import (
	"fmt"
	"time"
	"todo/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var client *gorm.DB

func Initmysql(listmodels ...interface{}) (*gorm.DB, error) {
	if client == nil {
		configvalue := config.GetConfig()
		count := 0
		newclient, err := gorm.Open(mysql.Open(configvalue.MySqlUri), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info),
		})
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
				newclient, err = gorm.Open(mysql.Open(configvalue.MySqlUri), &gorm.Config{})
			}
		}
		newclient.AutoMigrate(listmodels...)
		client = newclient
	}
	return client, nil
}
