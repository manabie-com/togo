package middleware

import (
	"github.com/gin-gonic/gin"
	pkg "github.com/manabie-com/togo/internal/pkg/utils"
)

func GinTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, _ := c.Cookie("Authorization")
		if len(tokenString) == 0 {
			tokenString = c.Request.Header.Get("Authorization")
		}
		// parse token
		claims, err := pkg.ParseToken(tokenString)
		if err != nil {
			return
		}

		userId, ok := claims["user_id"].(string)
		if !ok {
			return
		}
		// set user id
		c.Set("user_id", userId)
		c.Next()
	}
}
