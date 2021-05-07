package healthy

import (
	"github.com/gin-gonic/gin"
	"manabie-com/togo/global"
	"net/http"
	"strconv"
	"time"
)

func Info(c *gin.Context) {

	timezone, offset := time.Now().Zone()

	c.JSON(http.StatusOK, gin.H{
		"api_version": global.Config.ApiVersion,
		"server_time": timezone + " " + strconv.Itoa(offset),
		"up_time":     global.Uptime,
	})
}
