package dbcon

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/khoale193/togo/pkg/setting"
)

var sqlXDB *sqlx.DB
var gormDB *gorm.DB

// Setup initializes the database instance
func Setup() {
	setupSQLXDB()
	setupGormDB()
}

func GetGormDB() *gorm.DB {
	return gormDB
}

func GetSqlXDB() *sqlx.DB {
	return sqlXDB
}

func setupSQLXDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	sqlXDB, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
}

func setupGormDB() {
	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name)
	gormDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
}
