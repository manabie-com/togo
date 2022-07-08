package router

import (
	"encoding/json"
	"net/http"
	"todo/be/business"
	"todo/be/utils"

	"github.com/gorilla/mux"
)

type apiResult struct {
	Status  int    `json:"Status,omitempty"`
	Message string `json:"Message,omitempty"`
}

const (
	apiStatus_Success = 1
	apiStatus_Error   = 0
)

func initRouterApi(routerMux *mux.Router) {
	subRouter := routerMux.PathPrefix("/api/").Subrouter()
	subRouter.Use(apiMiddleware())
	subRouter.HandleFunc("/todo", apiTodo_add).Methods(http_POST)
}

func apiMiddleware() mux.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
			// do something
			h.ServeHTTP(response, request)
		})
	}
}

func apiTodo_add(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	user := getUserToken(response, request)
	var clientData map[string]string
	json.NewDecoder(request.Body).Decode(&clientData)
	err := business.Todo_add(request.Context(), user, clientData["Text"])
	if utils.IsError(err) {
		json.NewEncoder(response).Encode(apiResult{Status: apiStatus_Error, Message: err.Error()})
		return
	}
	json.NewEncoder(response).Encode(apiResult{Status: apiStatus_Success, Message: "Add todo task success"})
}
