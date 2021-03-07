package main

import (
	"database/sql"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/utils"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	sqlite "github.com/manabie-com/togo/internal/storages/sqlite"

	_ "github.com/mattn/go-sqlite3"

	"github.com/gorilla/mux"
)

func commonMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func authMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ok bool
		r, ok = transport.ValidToken(r)
		if !ok {
			utils.HttpResponseUnauthorized(w, "User is not valid!")
			return
		}
		next.ServeHTTP(w, r)
	})
}
func main() {
	var router = mux.NewRouter()
	router.Use(commonMiddleware)

	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	defer db.Close()

	taskStorage := sqlite.NewTaskLiteDBStorage(db)
	taskService := services.NewTaskService(taskStorage)
	taskTransport := transport.NewTaskTransport(taskService)

	userStorage := sqlite.NewUserLiteDBStorage(db)
	userService := services.NewUserService(userStorage)
	userTransport := transport.NewUserTransport(userService)

	router.HandleFunc("/login", userTransport.Login)
	taskAPI := router.PathPrefix("/tasks").Subrouter()
	taskAPI.Use(authMiddleware)
	taskAPI.HandleFunc("", taskTransport.ListTasks).Methods("GET")
	taskAPI.HandleFunc("", taskTransport.AddTasks).Methods("POST")

	http.ListenAndServe(":5050", router)

}
