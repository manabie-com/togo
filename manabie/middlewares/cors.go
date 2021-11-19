package middlewares

import (
	libs "manabie/manabie/helpers"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

// CORSMiddleware func
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

// ValidateToken func
func ValidateToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer libs.RecoverError(c)
		var (
			status = http.StatusOK
			msg    = ""
		)
		token := c.Request.Header.Get("Authorization")
		if token != "" {
			claims := make(jwt.MapClaims)
			t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("JWT_TOKEN")), nil
			})
			if err != nil {
				status = http.StatusUnauthorized
			}
			if !t.Valid {
				status = http.StatusUnauthorized
			}
			_, ok := claims["user_id"].(string)
			if !ok {
				status = http.StatusUnauthorized
			}
		} else {
			status = http.StatusUnauthorized
		}
		if status == http.StatusOK {
			c.Next()
		} else {
			if msg == "" {
				msg = "TOKEN_INVALID"
			}
			responseData := gin.H{
				"status": status,
				"msg":    msg,
			}
			c.JSON(status, responseData)
			c.Abort()
		}
	}
}
