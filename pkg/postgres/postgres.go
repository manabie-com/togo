package postgres

import (
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *PostgresConnect
	once     sync.Once
)

type PostgresConnect struct {
	db *gorm.DB
}

func GetMysqlConnInstance(dns string) *PostgresConnect {
	once.Do(func() {
		db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		instance = &PostgresConnect{db: db}
	})
	return instance
}

func NewPostgresConn(dns string) *gorm.DB {
	return GetMysqlConnInstance(dns).db
}

func (postgres *PostgresConnect) AutoMigrate(tables ...interface{}) error {
	for _, table := range tables {
		return postgres.db.AutoMigrate(&table)
	}
	return nil
}