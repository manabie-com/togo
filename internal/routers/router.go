package routers

import (
    "net/http"
	"sync"
	"database/sql"
	"log"

	"github.com/gorilla/mux"	
	"github.com/manabie-com/togo/internal/services"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
	_ "github.com/mattn/go-sqlite3"	

	"github.com/manabie-com/togo/internal/controllers"
	"github.com/manabie-com/togo/internal/storages"
	// "github.com/manabie-com/togo/internal/storages/sqllite"	
)

type IOhRouter interface {
	InitRouter() *mux.Router
}

type router struct{}

func (router *router) InitRouter() *mux.Router {


	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	sqliteHandler := &sqllite.SQLiteHandler{}
	sqliteHandler.DB = db
	userRepo := &storages.UserRepo{sqliteHandler}
	userService := &services.UserService{userRepo}
	userController := controllers.UserController{userService}

	quotaRepo := &storages.QuotaRepo{sqliteHandler}
	taskRepo := &storages.TaskRepo{sqliteHandler}
	taskService := &services.TaskService{taskRepo, quotaRepo}
	taskController := controllers.TaskController{taskService}

	r := mux.NewRouter()
	r.HandleFunc("/", HelloHandler).Methods("GET")
	r.HandleFunc("/login", userController.GetAuthToken).Methods("GET")
	r.HandleFunc("/tasks", taskController.ListTasks).Methods("GET")
	r.HandleFunc("/tasks", taskController.AddTask).Methods("POST")

	return r
}

var (
	m          *router
	routerOnce sync.Once
)

func OhRouter() IOhRouter {
	if m == nil {
		routerOnce.Do(func() {
			m = &router{}
		})
	}
	return m
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello! Who am I?\n"))
}