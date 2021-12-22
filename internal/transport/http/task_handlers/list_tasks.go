package task_handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/common/response"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/domain"
	"net/http"
)

func HttpListTasks(ctx *gin.Context) {
	userId, _ := ctx.Get(config.HEADER_USER_ID)

	taskDomain := domain.NewTaskDomain(ctx)
	tasks, err := taskDomain.GetlistTask(userId.(float64))

	if err == nil {
		ctx.JSON(http.StatusOK, response.Sucess(tasks))
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, response.Failure(err.Error()))
	}
}
