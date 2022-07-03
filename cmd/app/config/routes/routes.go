package routes

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/xrexonx/togo/cmd/app/config/environment"
	todoService "github.com/xrexonx/togo/internal/todo"
	"github.com/xrexonx/togo/internal/utils/response"
	"log"
	"net/http"
)

func Init() *mux.Router {

	router := mux.NewRouter()

	_apiPath := "/api/" + environment.GetValue("API_VERSION")

	router.HandleFunc(_apiPath+"/healthCheck", HealthCheckHandler)
	router.HandleFunc(_apiPath+"/todo", AddTodoHandler).Methods("POST")
	router.Use(logRequest)

	return router
}

// parseRequestBody parse request body using json new decoder instead on marshall/unmarshall
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
	newTodo, err := todoService.Add(todo)

	if err != nil {
		w.WriteHeader(http.StatusPreconditionFailed)
		response.Render(w, nil, http.StatusPreconditionFailed, err.Error(), "Precondition failed")
	} else {
		w.WriteHeader(http.StatusOK)
		response.Render(w, newTodo, http.StatusOK, "Successfully created todo", "Ok")
	}

}

// HealthCheckHandler A very simple health check.
func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response.Render(w, nil, http.StatusOK, "API is running", "Health Checked")
}

// logRequest request logger for debugging purposes
func logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Do debugging stuff here
		log.Println("api hit: " + r.RequestURI)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}
