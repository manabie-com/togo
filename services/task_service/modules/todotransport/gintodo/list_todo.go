package gintodo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	goservice "github.com/phathdt/libs/go-sdk"
	"github.com/phathdt/libs/go-sdk/sdkcm"
)

func ListTasks(sc goservice.ServiceContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, sdkcm.SimpleSuccessResponse("ok"))
	}
}
