package main

import (
	"lntvan166/togo/db"

	_ "github.com/lib/pq"
)

func main() {
	db.Connect()
}
