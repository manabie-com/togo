package main

import (
	"github.com/manabie-com/togo/internal/delivery/rest"
	"github.com/manabie-com/togo/internal/pkgs/clients"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/routers"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/tokens"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	dbConf := clients.PSQLConfig{
		DSN: "host=localhost user=togo password=ad34a$dg dbname=manabie_togo port=5432 sslmode=disable",
	}
	db, err := clients.InitPSQLDB(dbConf)
	if err != nil {
		log.Fatal("error opening db", err)
	}

	taskService := repositories.NewTaskRepo(db)
	userService := repositories.NewUserRepo(db)
	tokenService := tokens.NewTokenManager("wqGyEBBfPK9w3Lxw", userService)

	todoService := services.NewToDoService(taskService, tokenService)
	serializer := rest.NewSerializer(todoService)

	server := routers.NewServer(serializer)

	server.Run()
}
