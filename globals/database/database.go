package database

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ConfigDB struct {
	Username string
	Password string
	Host     string
	Port     string
	Dbname   string
}

var SQL *gorm.DB

func (c ConfigDB) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.Username,
		c.Password,
		c.Host,
		c.Port,
		c.Dbname,
	)
}

func InitDBConnection() {
	cfg := ConfigDB{
		Username:                 	os.Getenv("DB_USERNAME"),
		Password:               	os.Getenv("DB_PASSWORD"),
		Host:                 		os.Getenv("DB_HOSTNAME"),
		Dbname:               		os.Getenv("DB_NAME"),
		Port: 						os.Getenv("DB_PORT")}
	var err error
	SQL, err = gorm.Open(mysql.Open(cfg.GetDSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}