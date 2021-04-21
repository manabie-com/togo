package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/delivery"
	"github.com/manabie-com/togo/internal/respository"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/manabie-com/togo/internal/utils"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
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
		r, ok = delivery.ValidToken(r)
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

	taskRespository := respository.NewTaskLiteDBRespository(db)
	taskService := usecase.NewTaskService(taskRespository)
	taskDelivery := delivery.NewTaskDelivery(taskService)

	userRespository := respository.NewUserLiteDBRespository(db)
	userService := usecase.NewUserService(userRespository)
	userDelivery := delivery.NewUserDelivery(userService)

	router.HandleFunc("/login", userDelivery.Login)
	taskAPI := router.PathPrefix("/tasks").Subrouter()
	taskAPI.Use(authMiddleware)
	taskAPI.HandleFunc("", taskDelivery.ListTasks).Methods("GET")
	taskAPI.HandleFunc("", taskDelivery.AddTasks).Methods("POST")

	http.ListenAndServe(":5050", router)

}
