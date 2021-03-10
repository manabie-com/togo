package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/config"
	"log"
	"net/http"
)

func AuthMiddlewareHandler(handler func(writer http.ResponseWriter, request *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Let secure process the request. If it returns an error,
		// that indicates the request should not continue.
		token := req.Header.Get("Authorization")

		claims := make(jwt.MapClaims)
		t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(config.JWT_KEY), nil
		})
		if err != nil {
			log.Println(err)
			return
		}

		if !t.Valid {
			return
		}

		_, ok := claims["user_id"].(string)
		if !ok {
			return
		}

		handler(w, req)
	})
}
