package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/xrexonx/togo/cmd/app/config"
	todoService "github.com/xrexonx/togo/internal/todo"
	"gorm.io/gorm"
	"log"
	"net/http"
)

var DBInstance *gorm.DB

func Init(db *gorm.DB) *mux.Router {

	DBInstance = db
	router := mux.NewRouter()

	_apiPath := "/api/" + config.GetValue("API_VERSION")

	router.HandleFunc(_apiPath+"/health", HealthCheckHandler)
	router.HandleFunc(_apiPath+"/todo", AddTodoHandler).Methods("POST")
	router.Use(logRequest)

	return router
}

func parseRequestBody[T comparable](req *http.Request) T {
	reqBody := json.NewDecoder(req.Body)
	var data T
	err := reqBody.Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// AddTodoHandler Todo handler
func AddTodoHandler(w http.ResponseWriter, req *http.Request) {

	todo := parseRequestBody[todoService.Todo](req)
	newTodo, err := todoService.Add(DBInstance, todo)

	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		json.NewEncoder(w).Encode(todoService.Response{
			Status:  "Precondition Failed",
			Code:    http.StatusPreconditionFailed,
			Message: err.Error(),
		})
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(todoService.Response{
			Status:  "Success",
			Code:    http.StatusOK,
			Message: "Successfully created todo",
			Data:    newTodo,
		})
	}

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
