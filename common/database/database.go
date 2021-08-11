package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Database interface {
	GetDb() *gorm.DB
}

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DbName   string
}

func NewDB(setting *Config) Database {
	conn := fmt.Sprintf("host=%v port=%v dbname=%v user=%v password=%v sslmode=disable",
		setting.Host, setting.Port, setting.DbName, setting.UserName, setting.Password)
	db, err := gorm.Open("postgres", conn)
	if err != nil {
		panic(err)
	}
	return &database{
		db,
	}
}

type database struct {
	Db *gorm.DB
}

func (d *database) GetDb() *gorm.DB {
	return d.Db
}
