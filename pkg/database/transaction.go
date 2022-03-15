package database

import (
	"log"

	"gorm.io/gorm"
)

func BeginTransaction(db *gorm.DB) *gorm.DB {
	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	return tx
}

func ProcessTransaction(tx *gorm.DB, err error) {
	if err != nil {
		log.Printf(" ProcessTrasaction Error %s ", err.Error())
		tx.Rollback()
	} else {
		tx.Commit()
	}
}
