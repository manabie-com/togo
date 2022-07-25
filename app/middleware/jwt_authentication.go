package middleware

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/gorilla/mux"
	"github.com/huuthuan-nguyen/manabie/app/model"
	"github.com/huuthuan-nguyen/manabie/app/utils"
	"github.com/huuthuan-nguyen/manabie/config"
	"github.com/uptrace/bun"
	"log"
	"net/http"
	"strings"
)

// JWTAuthenticateMiddleware protect resources
func JWTAuthenticateMiddleware(config *config.Config, db bun.IDB) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			bearerToken := r.Header.Get("Authorization")
			if len(bearerToken) > 0 {
				bearerToken = strings.TrimPrefix(bearerToken, "Bearer ")
			} else {
				utils.PanicUnauthorized()
				return
			}

			payload := &Payload{}
			tkn, err := jwt.ParseWithClaims(bearerToken, payload, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method %v", token.Header["alg"])
				}
				return []byte(config.Server.Secret), nil
			})

			if err != nil || tkn == nil || !tkn.Valid {
				log.Println(err)
				utils.PanicUnauthorized()
				return
			}

			payload, ok := tkn.Claims.(*Payload)
			if !ok {
				utils.PanicUnauthorized()
				return
			}

			// prepare user
			user, err := model.FindOneUserByEmail(r.Context(), payload.Email, db)
			if err != nil {
				utils.PanicUnauthorized()
				return
			}

			ctx := context.WithValue(r.Context(), "user", &user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

type Payload struct {
	Email string `json:"email,omitempty"`
	jwt.StandardClaims
}
