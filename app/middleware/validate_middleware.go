package middleware

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/manabie/app/validator"
	"net/http"
)

// ValidateMiddleware add validator to request context
func ValidateMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			validate, _ := validator.NewValidator()
			ctx := context.WithValue(r.Context(), "validate", validate)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
