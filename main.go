package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/manabie-com/togo/config"
	"github.com/manabie-com/togo/internal/services"
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
	usecase "github.com/manabie-com/togo/internal/usecase"

	_ "github.com/lib/pq"
)

func main() {
	config, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("error connect to db:", err)
	}

	log.Println("Connected with DB!!")

	fmt.Printf("Server listening on port: %s\n", config.ServerAddress)
	DB := &postgres.PostgresDB{
		DB: db,
	}

	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerAddress), &services.ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		TaskUsecase: &usecase.TaskUsecase{
			Store: DB,
		},
		UserUsecase: &usecase.UserUsecase{
			Store: DB,
		},
	})
}
