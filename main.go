package main

import (
	"flag"
	"github.com/manabie-com/togo/internal/controller"
	"net/http"
)

var db string

func main() {
	flag.StringVar(&db, "db", "", "db run")
	flag.Parse()
	http.ListenAndServe(":5050", controller.NewToDoService(db))
}
