package database

import (
	"fmt"

	"github.com/manabie-com/togo/core/config"
	"github.com/pressly/goose/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	ManabieDB *gorm.DB
}

func New(d config.Config) *Database {
	c := d.Databases.PostgresDB
	connString := fmt.Sprintf("dbname=%v user=%v password=%v host=%v port=%v sslmode=%v", c.Database, c.Username, c.Password, c.Host, c.Port, c.SSLMode)
	db, err := gorm.Open(postgres.Open(connString), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect database successful !")
	return &Database{ManabieDB: db}
}

func (r *Database) Migrate() error {
	var err error
	if err = goose.SetDialect("postgres"); err != nil {
		return err
	}
	db, err := r.ManabieDB.DB()
	return goose.Up(db, "./script/database_script")
}
