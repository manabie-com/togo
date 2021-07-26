package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/handlers"
	"github.com/manabie-com/togo/internal/services/tasks"
	"github.com/manabie-com/togo/internal/services/users"
	"github.com/manabie-com/togo/internal/storages"
)

func index(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello there - Obiwan")
}

func main() {
	db, err := storages.Initialize("postgres", "postgres", "todo")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.DB.Close()

	userService := users.SetupNewService("wqGyEBBfPK9w3Lxw", db)
	taskService := tasks.SetupNewService(db)

	r := mux.NewRouter()
	//handlers
	n := negroni.Classic()
	handlers.MakeUserHandlers(r, *n, *userService)
	handlers.MakeTaskHandler(r, *n, *taskService, *userService)
	r.HandleFunc("/", index)

	n.UseHandler(r)
	n.Run(":5050")
}
