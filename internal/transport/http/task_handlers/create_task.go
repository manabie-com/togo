package task_handlers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/common/response"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/domain"
	"io/ioutil"
	"net/http"
)

func HttpCreateTask(ctx *gin.Context) {
	body, _ := ioutil.ReadAll(ctx.Request.Body)
	userId, _ := ctx.Get(config.HEADER_USER_ID)

	_data := make(map[string]string)
	err := json.Unmarshal(body, &_data)

	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, response.Failure(config.RESP_JSON_WRONG_FORMAT))
		return
	}

	taskDomain := domain.NewTaskDomain(ctx)
	message, status := taskDomain.CreateOneTask(userId.(float64), body)

	if status {
		ctx.JSON(http.StatusOK, response.Sucess(message))
	} else {
		ctx.JSON(http.StatusUnprocessableEntity, response.Failure(message))
	}
}
