package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Connect func
func Connect() *gorm.DB {
	for i := 1; i <= 3; i++ {
		db, err := gorm.Open(sqlite.Open("./manabie/databases/storages/data.db"), &gorm.Config{})
		if err == nil {
			return db
		} else {
			if i == 3 {
				panic(err.Error())
			}
		}
	}
	return nil
}
