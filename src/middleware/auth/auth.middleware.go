package AuthMiddleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Authorized blocks unauthorized requestrs
func Authorized(c *gin.Context) {
	_, exists := c.Get("user")
	if !exists {
		c.JSON(http.StatusForbidden, "Required Login")
		return
	}
}
