// Package middleware contains gin middleware
// Usage: router.Use(middleware.EnableCORS)
package middleware

import (
	"net/http"
	"strings"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	jwt "github.com/dgrijalva/jwt-go"
	// models "github.com/manabie-com/togo/internal/models"
)

type AuthHeader struct {
	IDToken string `header:"Authorization"`
}

var JWTKey string

// used to help extract validation errors
type invalidArgument struct {
    Field string `json:"field"`
    Value string `json:"value"`
    Tag   string `json:"tag"`
    Param string `json:"param"`
}


func ValidateToken(token string) (string, bool) {
	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return "", false
	}

	if !t.Valid {
		return "", false
	}

	userId, ok := claims["user_id"].(string)
	if !ok {
		return "",false
	}

	return userId, true
}

// ErrorHandler : ErrorHandler is a middleware to handle errors encountered during requests
func ErrorHandler(c *gin.Context) {
	// TODO: Handle it in a better way
	if len(c.Errors) > 0 {
		c.HTML(http.StatusBadRequest, "400", gin.H{
			"errors": c.Errors,
		})
	}

	c.Next()
}

// EnableCORS : Enable CORS Middleware in DEV Enviroment
func EnableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, Origin, Authorization, Accept, Client-Security-Token, Accept-Encoding, x-access-token")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func AuthUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := AuthHeader{}

		// bind Authorization Header to authHeader and check for validation errors
        if err := c.ShouldBindHeader(&authHeader); err != nil {
            c.JSON(500, gin.H{
                "error": err,
            })
            c.Abort()
            return
        }

		log.Println("authHeader: ", authHeader)

		if (AuthHeader{}) == authHeader  {
			c.JSON(http.StatusBadRequest, gin.H{
                "message": "Must provide Authorization header with format `Bearer {token}`",
            })
            c.Abort()
            return
		}

        idTokenHeader := strings.Split(authHeader.IDToken, "Bearer ")
		fmt.Println("idTokenHeader: " , idTokenHeader[1])
        if len(idTokenHeader) < 2 {
            c.JSON(http.StatusBadRequest, gin.H{
                "message": "Must provide Authorization header with format `Bearer {token}`",
            })
            c.Abort()
            return
        }

		// validate ID token here
        userId, isValid := ValidateToken(idTokenHeader[1])
		log.Println("user: ", userId)
		log.Println("isValid: ", isValid)
		
		if !isValid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthenticated",
			})
			c.Abort()
			return
		}

		userID, _ := uuid.Parse(userId)
		userCtx := make(map[string]interface{})
		userCtx["ID"] = userID

		c.Set("user", userCtx)
		c.Next()
	}
}
