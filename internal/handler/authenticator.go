package handler

import (
	"context"
	"errors"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

func NewAuthenticator(jwtKey string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			splitToken := strings.Split(token, "Bearer ")
			var err error

			if len(splitToken) == 2 {
				token = splitToken[1]

				claims := make(jwt.MapClaims)
				var t *jwt.Token
				t, err = jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
					return []byte(jwtKey), nil
				})

				if err == nil && t.Valid {
					userId, ok := claims["user_id"].(string)
					maxTodo, okMaxTodo := claims["max_todo"]
					if ok && okMaxTodo {
						ctx := context.WithValue(r.Context(), "userId", userId)
						ctx = context.WithValue(ctx, "maxTodo", int(maxTodo.(float64)))

						next.ServeHTTP(w, r.WithContext(ctx))
						return
					}
				}
			} else {
				err = errors.New("invalid bearer")
			}

			log.WithFields(log.Fields{
				"error": err,
				"token": token,
			}).Info("Invalid token")

			writeJsonRes(w, 401, errors.New("invalid token"))
		})
	}
}
