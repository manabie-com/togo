package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/models"
	pkg "github.com/manabie-com/togo/internal/pkg/utils"
	"github.com/manabie-com/togo/internal/services"
	"net/http"
	"strconv"
)

type Controller struct {
	utils       pkg.Utils
	taskService services.ToDoService
}

func New(
	taskService services.ToDoService,

) Controller {
	return Controller{
		taskService: taskService,
	}
}

func (c *Controller) GetQueryInt(key string, ctx *gin.Context) int {
	valueStr := ctx.Query(key)
	value, _ := strconv.Atoi(valueStr)
	return value
}

func (c *Controller) ErrorResponse(err error, ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, &models.Response{
		Code: "400",
		Msg:  err.Error(),
	})
}

func (c *Controller) Response(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, &models.Response{
		Code: "200",
		Msg:  "Successfully",
		Data: data,
	})
}
