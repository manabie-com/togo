package integration__tests

import (
	"log"
	"os"
	"testing"
	"togo/internal/pkg/domain/entities"
	db "togo/pkg/database"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

var (
	gormDB *gorm.DB
)

func TestMain(m *testing.M) {
	err := godotenv.Load(os.ExpandEnv("../.env"))
	if err != nil {
		log.Fatalf("Error getting env %v\n", err)
	}
	os.Exit(m.Run())
}

func database() error {
	dbConfig := db.DBConfig{
		Host: os.Getenv("DB_HOST_TEST"),
		Name: os.Getenv("DB_NAME_TEST"),
		User: os.Getenv("DB_USER_TEST"),
		Pass: os.Getenv("DB_PASS_TEST"),
		Port: os.Getenv("DB_PORT_TEST"),
	}
	db, err := db.NewDatabase(dbConfig)
	gormDB = db.DB
	return err
}

func Migrate() {
	gormDB.AutoMigrate(&entities.User{}, &entities.Todo{})
}
