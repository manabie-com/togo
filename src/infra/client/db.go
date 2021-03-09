package client

import (
	"log"
	"os"
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
		dbURL := os.Getenv("DATABASE_URL")
		if dbURL == "" {
			dbURL = "user=postgres password=nopass dbname=togo sslmode=disable"
		}
		conn, err := gorm.Open("postgres", dbURL)
		if err != nil {
			log.Fatal(err)
		}

		instance = &DB{
			PConn: conn,
		}
	})
	return instance
}
