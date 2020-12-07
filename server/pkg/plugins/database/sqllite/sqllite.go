package sqllite

import "database/sql"

func NewDatabase() *sql.DB{
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		panic("error opening db")
	}
	return db
}

