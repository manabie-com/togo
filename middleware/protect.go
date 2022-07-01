package middleware

import (
	"TOGO/untils"

	"context"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//"github.com/gorilla/context"
)

func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

		authHeader := strings.Split(r.Header.Get("Authorization"), "Bearer ")

		if len(authHeader) != 2 {
			untils.Error(rw, "token vaild", http.StatusBadRequest)
			return
		}

		jwtToken := authHeader[1]
		if jwtToken == "" {
			untils.Error(rw, "token vaild", http.StatusBadRequest)
			return
		}

		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(jwtToken, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_JWT")), nil
		})
		if err != nil {
			untils.Error(rw, "pls longin", http.StatusBadRequest)
			return
		}

		if !token.Valid {
			untils.Error(rw, "token invalid", http.StatusBadRequest)
			return
		}
		Role_User := claims["role"].(string)
		Id_User := claims["id"].(string)

		Role_Id := Role_User + " " + Id_User
		ctx := context.WithValue(r.Context(), "Role_Id", Role_Id)
		next.ServeHTTP(rw, r.WithContext(ctx))
	})
}

func CreateToken(id primitive.ObjectID, role string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["role"] = role
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Hour * 1200).Unix() //Token expires after 12h
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("SECRET_JWT")))
}
