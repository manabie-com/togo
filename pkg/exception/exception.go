package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// ThrownError return bad request with err != nil
func ThrownError(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, err.Error())
}
