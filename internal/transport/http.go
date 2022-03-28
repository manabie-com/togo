package transport

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHttpServer(
	jwtKey string,
	authHandler AuthHandler,
	taskHandler TaskHandler,
) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost)
	router.HandleFunc("/tasks", Adapt(http.HandlerFunc(taskHandler.GetList), WithAuth(jwtKey)).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/task", Adapt(http.HandlerFunc(taskHandler.AddTask), WithAuth(jwtKey)).ServeHTTP).Methods(http.MethodPost)

	return router
}

type Adapter func(http.Handler) http.Handler

func Adapt(handler http.Handler, adapters ...Adapter) http.Handler {
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}
	return handler
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	data := map[string]interface{}{
		"data": object,
	}
	json.NewEncoder(writer).Encode(data)
}
