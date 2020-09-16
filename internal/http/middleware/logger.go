package middleware

import (
  "log"
  "net/http"
)

type Logger struct{}

func (logger Logger) MethodAndPath(next http.Handler) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    log.Printf("[http::Server::ServeHTTP - %s %s]\n", r.Method, r.URL.Path)
    next.ServeHTTP(w, r)
  }
}
