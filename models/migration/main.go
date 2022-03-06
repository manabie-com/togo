package migration

import (
	"fmt"
	"log"
	"reflect"
	"togo/models"

	"gorm.io/gorm"

	"togo/models/dbcon"
)

func Migrate() {
	gormDB := dbcon.GetGormDB()
	gormDB.AutoMigrate(&models.Member{})
	gormDB.AutoMigrate(&models.Task{})

	gormDB.AutoMigrate(&Migration{})
	_ = migrate(1)
}

func migrate(version int) error {
	gormDB := dbcon.GetGormDB()
	var migrate Migration
	err := gormDB.Order("id desc").Where("id <= ?", version).Take(&migrate).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}
	migrateNoFrom := 1
	if migrate.ID != 0 {
		if int(migrate.ID) == version {
			migrateNoFrom = version + 1
		} else {
			migrateNoFrom = int(migrate.ID) + 1
		}
	}
	for i := migrateNoFrom; i <= version; i++ {
		callableName := fmt.Sprintf("Migrate%dUp", i)
		callable := reflect.ValueOf(&Migration{}).MethodByName(callableName)
		if !callable.IsValid() {
			log.Fatalf("\nFailed to find Registry method with name \"%s\".", callableName)
		}
		callable.Call([]reflect.Value{})
		migrate := Migration{
			Migration: callableName,
			Batch:     1,
		}
		_ = gormDB.Create(&migrate).Error
	}
	return nil
}
