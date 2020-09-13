package http

import "net/http"

type AuthMw interface {
  SetUser(next http.Handler) http.HandlerFunc
  RequireUser(next http.Handler) http.HandlerFunc
}
