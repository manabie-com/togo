package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/api/handler"
)

// CheckUserId check user id if exist on request. (Intended for token based user/ session)
func CheckUserId() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.Header.Get(handler.UserIdHeader) == "" {
			//Return bad request for now instead of 401
			context.AbortWithStatusJSON(http.StatusBadRequest, map[string]interface{}{
				"message": "Please identify user",
			})
		} else {
			context.Next()
		}

	}
}
