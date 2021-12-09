package mysql

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	instance *MysqlConnect
	once     sync.Once
)

type MysqlConnect struct {
	db *gorm.DB
}

func GetMysqlConnInstance(dns string) *MysqlConnect {
	once.Do(func() {
		db, err := gorm.Open(mysql.Open(dns), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance = &MysqlConnect{db: db}
	})
	return instance
}

func NewMysqlConn(dns string) *gorm.DB {
	return GetMysqlConnInstance(dns).db
}

func (mysql *MysqlConnect) AutoMigrate(tables ...interface{}) error {
	for _, table := range tables {
		return mysql.db.AutoMigrate(&table)
	}
	return nil
}