package main

import (
	"log"
	"net/http"

	"github.com/manabie-com/togo/define"
	"github.com/manabie-com/togo/driver"
	"github.com/manabie-com/togo/router"
)

func main() {

	db, err := driver.GetConnection()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	http.ListenAndServe(":5050", &router.Router{
		JWTKey: define.JWTKey,
		Conn:   db.Conn,
	})
}
