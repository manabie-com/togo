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

var (
	drivers = []Driver{
		{Name: "Jimmy Johnson", License: "ABC123"},
		{Name: "Howard Hills", License: "XYZ789"},
		{Name: "Craig Colbin", License: "DEF333"},
	}
	cars = []Car{
		{Year: 2000, Make: "Toyota", ModelName: "Tundra", DriverID: 1},
		{Year: 2001, Make: "Honda", ModelName: "Accord", DriverID: 1},
		{Year: 2002, Make: "Nissan", ModelName: "Sentra", DriverID: 2},
		{Year: 2003, Make: "Ford", ModelName: "F-150", DriverID: 3},
	}
)
type Driver struct {
	gorm.Model
	Name    string
	License string
	Cars    []Car
}

type Car struct {
	gorm.Model
	Year      int
	Make      string
	ModelName string
	DriverID  int
}
func (a *App) Initialize() {
	var err error
	a.DB, err = config.GetPostgersDB()

	if err != nil {
		log.Fatal(err)
	}
	a.DB.AutoMigrate(&Driver{})
	a.DB.AutoMigrate(&Car{})

	for index := range cars {
		a.DB.Create(&cars[index])
	}

	for index := range drivers {
		a.DB.Create(&drivers[index])
	}

	var car Car
	a.DB.First(&car)

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

