package config

import (
	"fmt"
	"log"
	"os"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var once sync.Once

func GetDBUrl() string {
	return fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=%s password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL"),
		os.Getenv("DB_PASS"))
}

func GetDBDriver() string {
	return os.Getenv("DB_DRIVER")
}

func Database() *gorm.DB {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(postgres.Open(GetDBUrl()), &gorm.Config{})
		if err != nil {
			log.Panic(err)
		}

		log.Println("Database Connected")
	})

	return DB
}

func DBMock() *gorm.DB {
	once.Do(func() {
		var err error
		DB, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
		if err != nil {
			log.Panic(err)
		}

		log.Println("Database Mock Connected")

		err = DB.Exec(`
			drop table if exists tasks;

			create table if not exists tasks (
			id serial primary key,
			content text,
			user_id varchar(36),
			status smallint default 1,
			created_at timestamp without time zone default current_timestamp,
			updated_at timestamp without time zone default current_timestamp,
			deleted_at timestamp without time zone default NULL
		);
		`).Error
		if err != nil {
			log.Panic(err)
		}
	})
	return DB
}
