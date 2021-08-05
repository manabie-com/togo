package test

import (
	"github.com/jinzhu/gorm"
	db2 "github.com/manabie-com/togo/pkg/db"
)

func GetTestDb() *gorm.DB {
	connStr := `host=localhost port=5432 user=admin password=admin dbname=test sslmode=disable binary_parameters=yes`
	db, err := db2.NewDatabase("postgres", connStr, 10, 2)
	if err != nil {
		panic("no db connection")
	}
	return db
}
