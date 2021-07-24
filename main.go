package main

import (
	_ "github.com/mattn/go-sqlite3"
	"github.com/surw/togo/internal/services"
)

func main() {
	service := services.NewToDoService()
	service.Serve(5050, services.NewRouter())
}
