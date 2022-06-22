package db

import (
	"database/sql"
	"fmt"
)

const ()

func GetPostgresConnectionString() string {
	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "handsome2022"
		dbname   = "togo"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	return psqlInfo
}

func Connect() {
	psqlInfo := GetPostgresConnectionString()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
}
