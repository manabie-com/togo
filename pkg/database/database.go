package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/gorm/logger"

	"github.com/manabie-com/togo/core/config"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	ManabieDB     *gorm.DB
	TestManabieDB *gorm.DB
}

func New(d config.DBConfig) *Database {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Disable color
		},
	)

	var (
		db, testDB *gorm.DB
		err        error
	)

	if d.PostgresDB != nil {
		c := d.PostgresDB
		connString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", c.Database, c.Username, c.Password, c.Host, c.Port, c.SSLMode)
		db, err = gorm.Open(postgres.Open(connString), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 newLogger,
		})
		if err != nil {
			panic(err)
		}
		fmt.Println("Connect database successful !")
	}

	if d.Test_PostgresDB != nil {
		t := d.Test_PostgresDB
		testConnString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", t.Database, t.Username, t.Password, t.Host, t.Port, t.SSLMode)
		testDB, err = gorm.Open(postgres.Open(testConnString), &gorm.Config{
			SkipDefaultTransaction: true,
			Logger:                 newLogger,
		})
		if err != nil {
			panic(err)
		}
	}

	return &Database{
		ManabieDB:     db,
		TestManabieDB: testDB,
	}
}

func (r *Database) Migrate() error {
	var err error
	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}
	db, err := r.ManabieDB.DB()
	return goose.Up(db, "./script/database_script")
}

func (r *Database) MigrateTestDB() error {
	var err error
	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}
	db, err := r.TestManabieDB.DB()
	return goose.Up(db, "./script/test_database_script")
}
