package jwt

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"togo/pkg/e"
)

// JWT is jwt middleware
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := e.SUCCESS
		token := c.Request.Header.Get(e.UserAuth)
		if token == "" {
			code = e.INVALID_PARAMS
		} else {

		}
		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
