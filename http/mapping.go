package http

import (
	"net/http"
)

func RegisterService(handler http.Handler) {
	http.Handle("/api/", handler)
}
