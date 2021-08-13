package main

import (
	"github.com/manabie-com/togo/internal/controller"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/jinzhu/gorm"
	"github.com/rs/cors"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)
import "github.com/manabie-com/togo/internal/config"

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize() {
	var err error
	a.DB, err = config.GetPostgersDB()

	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

    a.initializeRoutes()
}

func (a *App) Run(addr string) {
	//handler := cors.Default().Handler(a.Router)
	corsWrapper := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
		AllowedHeaders: []string{"Content-Type", "Origin", "Accept", "*"},
	})
	log.Fatal(http.ListenAndServe(addr, corsWrapper.Handler(a.Router)))
}

func (a *App) initializeRoutes() {
	auth := controller.NewAuthController(a.DB)
	task := controller.NewTaskController(a.DB)
	a.Router.Handle("/login", http.HandlerFunc(auth.GetAuthToken)).Methods("GET")
	a.Router.Handle("/task", config.Middleware(http.HandlerFunc(task.ListTasks))).Methods("GET")
	a.Router.Handle("/task", config.Middleware(http.HandlerFunc(task.AddTask))).Methods("POST")
}

