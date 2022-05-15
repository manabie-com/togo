package database

import (
	"embed"
	"log"
	"os"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/pressly/goose/v3"
)

var (
	db *gorm.DB

	//go:embed migrations/*.sql
	f embed.FS
)

const (
	defaultMaxIdleConns = 1
	defaultMaxOpenConns = 1
)

func Init() {
	var err error
	db, err = gorm.Open("postgres", os.Getenv("DB_DSN"))
	if err != nil {
		log.Fatal("connect db error\n", err)
	}

	// Set log mode
	if os.Getenv("LOG_MODE") == "true" {
		db.LogMode(true)
	}

	// Set max idle conns
	maxIdleConns := os.Getenv("MAX_IDLE_CONNS")
	intMaxIdleConns, err := strconv.Atoi(maxIdleConns)
	if err != nil {
		log.Printf("get Max Idle Conns error, use default")
		intMaxIdleConns = defaultMaxIdleConns
	}
	db.DB().SetMaxIdleConns(intMaxIdleConns)

	// Set max open conns
	maxOpenConns := os.Getenv("MAX_OPEN_CONNS")
	intMaxOpenConns, err := strconv.Atoi(maxOpenConns)
	if err != nil {
		log.Printf("get Max Open Conns error, use default")
		intMaxOpenConns = defaultMaxOpenConns
	}
	db.DB().SetMaxOpenConns(intMaxOpenConns)

	// migrate
	err = migrate()
	if err != nil {
		panic("migration error")
	}
}

// migrate call migration up everytime before start server
func migrate() error {
	goose.SetBaseFS(f)
	err := goose.SetDialect("postgres")
	if err != nil {
		return err
	}

	return goose.Up(db.DB(), "migrations")
}

// DB return db
func DB() *gorm.DB {
	return db
}

// Close close db after after the server is down
func Close() {
	if db == nil {
		return
	}
	err := db.Close()
	if err != nil {
		log.Print("close db error\n", err)
	}
}
