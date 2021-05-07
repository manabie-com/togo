package healthy

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Healthy(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"detail": "I'm strong and healthy!",
	})
}
