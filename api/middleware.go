package api

import (
	"net/http"
	"togo/app"
	"togo/validation"

	"github.com/gin-gonic/gin"
)

func authenMiddleware(c *gin.Context) {
	token, err := c.Cookie(app.COOKIE_ACCESS_TOKEN)
	if err != nil {
		requestToken := c.Request.Header.Get("Authorization")
		token = requestToken
	}

	// verify token
	claims, err := validation.VerifyToken(token)
	if err != nil {
		if err.Error() == validation.ErrTokenExpiredMsg {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": validation.ErrTokenExpiredMsg,
			})
			return
		}

		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	//set userID into header for next
	c.Request.Header.Set("user_id", claims.UID)
	c.Next()
}
