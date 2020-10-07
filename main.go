package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	app "github.com/manabie-com/togo/internal"
)

func main() {
	app.Run()
}
