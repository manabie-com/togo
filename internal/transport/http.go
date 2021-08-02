package transport

import (
	"encoding/json"
	"net/http"
)

func NewHttpServer(
	jwtKey string,
	authHandler AuthHandler,
	taskHandler TaskHandler,
) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/login", authHandler.Login)
	mux.HandleFunc("/register", authHandler.Register)
	mux.Handle("/tasks", WithCors(
		WithAuth(
			taskRouter(taskHandler),
			jwtKey,
		),
	),
	)

	return mux
}

func taskRouter(taskHandler TaskHandler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			taskHandler.GetList(w, r)
		case http.MethodPost:
			taskHandler.AddTask(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}

func responseWithJson(writer http.ResponseWriter, status int, object interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(status)
	data := map[string]interface{}{
		"data": object,
	}
	json.NewEncoder(writer).Encode(data)
}
