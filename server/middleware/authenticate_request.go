package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"togo/common/key"
	"togo/common/response"

	"github.com/golang-jwt/jwt"
)

type ctxkey int

const (
	keyPrincipalID ctxkey = iota
)

func AuthenticateRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(response.ErrorResponse{
					Status:  "fail",
					Code:    http.StatusUnauthorized,
					Message: err.Error(),
				})
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.ErrorResponse{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		tokenStr := cookie.Value

		claims := &key.Claims{}

		tkn, err := jwt.ParseWithClaims(tokenStr, claims,
			func(t *jwt.Token) (interface{}, error) {
				return key.JwtKey, nil
			})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(response.ErrorResponse{
					Status:  "fail",
					Code:    http.StatusUnauthorized,
					Message: err.Error(),
				})
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(response.ErrorResponse{
				Status:  "fail",
				Code:    http.StatusBadRequest,
				Message: err.Error(),
			})
			return
		}

		if !tkn.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(response.ErrorResponse{
				Status:  "fail",
				Code:    http.StatusUnauthorized,
				Message: err.Error(),
			})
			return
		}

		ctx := context.WithValue(r.Context(), keyPrincipalID, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
