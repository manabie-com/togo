package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/shared"
	apperror "github.com/manabie-com/togo/shared/app_error"
	"net/http"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json; charset=UTF-8")

				if appErr, ok := err.(*apperror.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					return
				}

				appErr := shared.ErrInternal(err.(error))
				c.AbortWithStatusJSON(http.StatusInternalServerError, appErr)
			}
		}()
		c.Next()
	}
}
