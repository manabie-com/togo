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
		r, ok = usecase.ValidToken(r)
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
	taskService := delivery.NewTaskService(taskRespository)
	taskUsecase := usecase.NewTaskUsecase(taskService)

	userRespository := respository.NewUserLiteDBRespository(db)
	userService := delivery.NewUserService(userRespository)
	userUsecase := usecase.NewUserUsecase(userService)

	router.HandleFunc("/login", userUsecase.Login)
	taskAPI := router.PathPrefix("/tasks").Subrouter()
	taskAPI.Use(authMiddleware)
	taskAPI.HandleFunc("", taskUsecase.ListTasks).Methods("GET")
	taskAPI.HandleFunc("", taskUsecase.AddTasks).Methods("POST")

	http.ListenAndServe(":5050", router)

}
