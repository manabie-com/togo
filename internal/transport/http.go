package transport

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func NewHttpServer(
	jwtKey string,
	authHandler AuthHandler,
) http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/login", authHandler.Login).Methods(http.MethodPost)
	router.HandleFunc("/register", authHandler.Register).Methods(http.MethodPost)
	// router.HandleFunc("/tasks",Adapt(task))
	// mux.Handle("/tasks", WithCors(
	// 	WithAuth(
	// 		taskRouter(taskHandler),
	// 		jwtKey,
	// 	),
	// ),
	// )

	return router
}

type Adapter func(http.Handler) http.Handler

func Adapt(handler http.Handler, adapters ...Adapter) http.Handler {
	for i := len(adapters); i > 0; i-- {
		handler = adapters[i-1](handler)
	}
	return handler
}

// func taskRouter(taskHandler TaskHandler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		switch r.Method {
// 		case http.MethodGet:
// 			taskHandler.GetList(w, r)
// 		case http.MethodPost:
// 			taskHandler.AddTask(w, r)
// 		default:
// 			w.WriteHeader(http.StatusMethodNotAllowed)
// 		}
// 	})
// }

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	data := map[string]interface{}{
		"data": object,
	}
	json.NewEncoder(writer).Encode(data)
}
