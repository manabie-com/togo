package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	a := App{}

	a.Initialize()

	a.Run(":5050")
}
