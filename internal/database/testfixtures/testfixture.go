package testfixture

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/manabie-com/togo/internal/env"

	"github.com/jinzhu/gorm"
	"gopkg.in/testfixtures.v2"
)

func init() {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USERNAME", "postgres")
	os.Setenv("DB_PASSWORD", "123456")
	os.Setenv("DB_SSL_MODE", "disable")
	os.Setenv("DB_NAME_TEST", "togo_test")
}

func SetupRepo(fixturePath string) *gorm.DB {
	connectionStr := fmt.Sprintf("host=%s port=%s sslmode=%s user=%s password=%s dbname=%s",
		env.GetDBHost(), env.GetDBPort(), env.GetDBSSLMode(), env.GetDBUserName(), env.GetDBPassword(), env.GetDBTest())

	db, err := sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
	}
	gormDB, _ := gorm.Open("postgres", db)

	fixtures, err := testfixtures.NewFolder(db, &testfixtures.PostgreSQL{}, fixturePath)
	if err != nil {
		log.Fatal(err)
	}
	if err := fixtures.Load(); err != nil {
		log.Fatal(err)
	}

	gormDB.DB()
	gormDB.DB().Ping()

	gormDB.SingularTable(false)
	gormDB.LogMode(true)

	return gormDB
}
