package main

import (
	"log"

	"github.com/manabie-com/backend/app"
	"github.com/manabie-com/backend/controller/tasks"
	"github.com/manabie-com/backend/repository"
	taskservice "github.com/manabie-com/backend/services/task"
)

func main() {

	repo := repository.NewRepository()
	serv := taskservice.NewTaskService(repo)
	controller := tasks.NewTaskController(serv)
	server, err := app.NewServer(controller)
	if err != nil {
		log.Fatal("cannot create server:", err)
	}

	errStart := server.Start("0.0.0.0:8080")
	if errStart != nil {
		log.Fatal("cannot create server:", err)
	}

}
