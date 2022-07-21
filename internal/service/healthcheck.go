package service

import (
	"io"
	"net/http"
)

func healthcheck(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok")
}

func (s *Service) Health(w http.ResponseWriter, req *http.Request) {
	healthcheck(w, req)
}
