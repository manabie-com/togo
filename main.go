package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages/postgres"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbconnecter, err := config.GetDBConnecter()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := dbconnecter.Connect()
	if err != nil {
		log.Fatal("error opening db", err)
	}

	log.Println("connect database successfully")

	switch v := dbconnecter.(type) {
	case *config.Postgres:
		log.Printf("use database %v\n", v)
		http.ListenAndServe(":5050", &services.ToDoService{
			JWTKey: "wqGyEBBfPK9w3Lxw",
			Store: &postgres.PqDB{
				DB: db,
			},
		})
	case *config.Sqlite:
		log.Printf("use database %v\n", v)
		http.ListenAndServe(":5050", &services.ToDoService{
			JWTKey: "wqGyEBBfPK9w3Lxw",
			Store: &sqllite.LiteDB{
				DB: db,
			},
		})
	}

}
