package servehttp

import "net/http"

type IAPIHandler interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}
