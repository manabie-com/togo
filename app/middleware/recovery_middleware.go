package middleware

import (
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/manabie/app/utils"
	"net/http"
)

// RecoveringMiddleware wrap the gorilla Recovery handler
func RecoveringMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// recover when panic happening
			defer func() {
				if err := recover(); err != nil {
					// Handling the error
					if e, ok := err.(utils.Error); ok {
						statusCode := e.StatusCode
						w.WriteHeader(statusCode)
						_, _ = w.Write([]byte(e.Error()))
					} else {
						w.WriteHeader(http.StatusInternalServerError)
						_, _ = w.Write([]byte("Something went wrong."))
					}
				}
			}()

			next.ServeHTTP(w, r)
		})
	}
}
