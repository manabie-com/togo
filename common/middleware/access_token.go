package middleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

const ContextKeyUserId = "ctx_user_id"
const UserIdClaimName = "user_id"

func AccessToken(jwtKey string) gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.Request.Header.Get("Authorization")
		token = strings.TrimPrefix(token, "Bearer ")
		claims := make(jwt.MapClaims)
		t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		})
		if err != nil {
			log.Println(err)
			context.Writer.WriteHeader(http.StatusUnauthorized)
			context.Abort()
		}

		if !t.Valid {
			context.Writer.WriteHeader(http.StatusUnauthorized)
			context.Abort()
		}

		id, ok := claims[UserIdClaimName].(string)
		if !ok {
			context.Writer.WriteHeader(http.StatusUnauthorized)
			context.Abort()
		}
		context.Set(ContextKeyUserId, id)
	}
}

func UserIdFromAccessToken(ctx context.Context) string {
	return ctx.Value(ContextKeyUserId).(string)
}
