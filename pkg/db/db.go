package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func NewDatabase(extension, conStr string, OpenConnect, IdleCon int) (*gorm.DB, error) {
	db, err := gorm.Open(extension, conStr)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(OpenConnect)
	db.DB().SetMaxIdleConns(IdleCon)
	return db, nil
}

func Migration(db *gorm.DB, models []interface{}) {
	for _, m := range models {
		db.AutoMigrate(m)
	}
}
