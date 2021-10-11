package main

import (
	"log"
	"net/http"

	"github.com/cuongtop4598/togo-interview/togo/internal/helper"
	"github.com/cuongtop4598/togo-interview/togo/internal/services"
	"github.com/cuongtop4598/togo-interview/togo/internal/storages/postgres"
)

func main() {
	err := helper.AutoBindConfig("./internal/storages/postgres/config.yml")
	if err != nil {
		panic(err)
	}
	db, err := postgres.NewDBManager()
	if err != nil {
		panic(err)
	}
	log.Println("listion and serve on :5510")
	http.ListenAndServe(":5510", &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store:  db.(*postgres.DBmanager),
	})

}
