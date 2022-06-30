package repository

import (
	"database/sql"
	"fmt"
	"lntvan166/togo/internal/config"
)

func GetPostgresConnectionString() string {
	host := config.DB_HOST
	port := config.DB_PORT
	user := config.DB_USER
	password := config.DB_PASSWORD
	dbname := config.DB_NAME

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}

func Connect() *sql.DB {
	var (
		err error
		db  *sql.DB
	)
	psqlInfo := config.DATABASE_URL
	if psqlInfo == "" {
		psqlInfo = GetPostgresConnectionString()
	}

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}