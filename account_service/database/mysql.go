package database

import (
	"database/sql"
	"os"
)

//DBInfo ...
type DBInfo struct {
	username string
	password string
	url      string
	port     string
	database string
}

func (db *DBInfo) setDB() {
	db.username = os.Getenv("MYSQL_USERNAME")
	db.password = os.Getenv("MYSQL_PASSWORD")
	db.url = os.Getenv("MYSQL_URL")
	db.port = os.Getenv("MYSQL_PORT")
	db.database = os.Getenv("MYSQL_DATABASE")
}

//GetDB ...
func (db *DBInfo) GetDB() (*sql.DB, error) {
	db.setDB()
	return sql.Open("mysql", db.username+":"+db.password+"@tcp"+"("+db.url+":"+db.port+")/"+db.database)
}
