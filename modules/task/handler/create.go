package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"togo/pkg/app"
	"togo/pkg/e"
)

func CreateTask(c *gin.Context) {
	var appG = app.Gin{C: c}
	appG.Response(c, http.StatusOK, e.GetMsg(e.SUCCESS), e.SUCCESS, nil)
}
