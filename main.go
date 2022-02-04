package main

import (
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	api "github.com/roandayne/togo/api"
)

func cors(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "*")
		w.Header().Set("Content-type", "application/json")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func main() {
	// API routes
	r := mux.NewRouter()
	r.Use(cors)
	r.HandleFunc("/api/tasks", api.CreateTask).Methods("POST")
	http.ListenAndServe(":8080", r)
}
