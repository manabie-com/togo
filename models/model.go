package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

var db *gorm.DB
var err error

func InitFromSQLLite(db_path_str string) (func(), error) {
	db, closeFunc, err := NewGormDB(db_path_str)
	if err != nil {
		panic(err)
	}
	err = migrateTable(db)
	migrateData()
	return closeFunc, err
}

func NewGormDB(db_path_str string) (*gorm.DB, func(), error) {
	db, err = gorm.Open("sqlite3", db_path_str)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return defaultTableName
	}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	cleanFunc := func() {
		err := db.Close()
		if err != nil {
			log.Fatal()
		}
	}
	return db, cleanFunc, err
}

func migrateTable(db *gorm.DB) error {
	return db.AutoMigrate(
		new(User),
		new(Task),
	).Error
}

func migrateData()  {
	(&User{}).MigrateData()
}