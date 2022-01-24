package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"togo/common"
)



func Recovery () gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			c.Header("Content-Type", "application/json")

			if err := recover(); err != nil {
				c.AbortWithStatusJSON(common.SERVER_ERROR_STATUS, common.ResponseError(common.SERVER_ERROR_STATUS, fmt.Sprintf("%v", err)))
			}
		}()

		c.Next()
	}
}