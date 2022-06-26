package middlewares

import "net/http"

type Middleware interface {
	WithCors() func(http.Handler) http.Handler
}

type middleware struct {
}

func NewMiddleware() Middleware {
	return middleware{}
}
