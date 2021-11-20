package db

import (
	"errors"
	"mini_project/db/model"
	"mini_project/db/mysql_driver"

	"gorm.io/gorm"
)

var (
	InvalidDBType = errors.New("invalid db type")
)

// Get database
func GetDatabase(dbUrl map[string]string) (model.DatabaseAPI, error) {
	switch dbUrl["Type"] {
	case "mysql":
		return mysql_driver.GetConnection(dbUrl), nil
	default:
		return nil, InvalidDBType
	}
}

func GetMigrator(dbUrl map[string]string) (gorm.Migrator, error) {
	switch dbUrl["Type"] {
	case "mysql":
		return mysql_driver.GetMigrator(dbUrl), nil
	default:
		return nil, InvalidDBType
	}
}

func CreateDatabase(dbUrl map[string]string) error {
	switch dbUrl["Type"] {
	case "mysql":
		mysql_driver.New(dbUrl)
		return nil
	default:
		return InvalidDBType
	}
}

// cleanup data
// Note: use this for developing only
func PurgeDB(dbUrl map[string]string) {
	db, _ := GetMigrator(dbUrl)

	if db.HasTable(&model.User{}) {
		db.DropTable(&model.User{})
	}

}
