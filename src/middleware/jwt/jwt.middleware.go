package JwtMiddleware

import (
	"strings"
	jwt "togo/src/common/jwt"
	"togo/src/common/types"
	Users "togo/src/modules/users"

	"github.com/gin-gonic/gin"
)

// JWTMiddleware parses JWT token from cookie and stores data and expires date to the context
// JWT Token can be passed as cookie, or Authorization header
func ParseUserContext() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString, err := c.Cookie("token")
		// failed to read cookie
		if err != nil {
			// try reading HTTP Header
			authorization := c.Request.Header.Get("Authorization")
			if authorization == "" {
				c.Next()
				return
			}
			sp := strings.Split(authorization, "Bearer ")
			// invalid token
			if len(sp) < 1 {
				c.Next()
				return
			}
			tokenString = sp[1]
		}

		tokenData, err := jwt.ValidateToken(tokenString)
		if err != nil {
			c.Next()
			return
		}

		var user Users.User
		user.Read(tokenData["user"].(types.JSON))

		c.Set("user", user)
		c.Set("token_expire", tokenData["exp"])
		c.Next()
	}
}
