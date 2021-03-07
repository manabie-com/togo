package client

import (
	"log"
	"sync"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type DB struct {
	PConn *gorm.DB
}

var (
	instance *DB
	once     sync.Once
)

func NewDB() *DB {
	once.Do(func() {
		conn, err := gorm.Open("postgres", "user=postgres password=nopass dbname=togo sslmode=disable")
		if err != nil {
			log.Fatal(err)
		}

		instance = &DB{
			PConn: conn,
		}
	})
	return instance
}
