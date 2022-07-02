package routes

import (
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
)

const _apiPath = "/api/v1"

var DBInstance *gorm.DB

func Init(db *gorm.DB) *mux.Router {

	DBInstance = db
	router := mux.NewRouter()

	router.HandleFunc(_apiPath+"/health", HealthCheckHandler)
	router.HandleFunc(_apiPath+"/todo", AddTodoHandler).Methods("POST")
	router.Use(logRequest)

	return router
}

// AddTodoHandler Todo handler
func AddTodoHandler(w http.ResponseWriter, req *http.Request) {

}

// HealthCheckHandler A very simple health check.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Health Checked"))
}

// logRequest request logger for debugging purposes
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do debugging stuff here
		log.Println(r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
