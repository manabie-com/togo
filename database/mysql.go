package database

import (
	"fmt"
	"log"
	"os"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/luongdn/togo/config"
	"github.com/luongdn/togo/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SQLdb *gorm.DB

func ConnectMySQL() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Config.MySQL.User,
		config.Config.MySQL.Pass,
		config.Config.MySQL.Host,
		config.Config.MySQL.Port,
		config.Config.MySQL.Name,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("MySQL connected")
	db.AutoMigrate(&models.User{}, &models.Task{}, &models.Rule{})
	SQLdb = db
}

func Seed() {
	user := models.User{}
	result := SQLdb.First(&user)
	if result.RowsAffected > 0 {
		return
	}

	user = models.User{
		ID:       "9725cc63-4e92-4893-a6b2-216617f3a5dd",
		Username: "user",
	}
	SQLdb.Create(&user)

	rule := models.Rule{
		UserID:          user.ID,
		Action:          models.TaskCreate,
		Unit:            models.Day,
		RequestsPerUnit: 3,
	}
	SQLdb.Create(&rule)
}

// Open a mock database connection used for testing
func OpenMockDBConn() (*gorm.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("error while opening a stub database connection: '%s'", err)
	}

	SQLdb, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	if err != nil {
		log.Fatalf("error connecting database: '%s'", err)
	}

	return SQLdb, mock
}
