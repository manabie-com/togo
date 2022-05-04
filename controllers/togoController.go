package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/manabie-com/togo/models"

	"github.com/manabie-com/togo/factories"
)

func AddTogoTask(context *gin.Context) {
	var newtogo models.Togo

	if err := context.BindJSON(&newtogo); err != nil {
		context.JSON(http.StatusBadRequest, err.Error())
		return
	}

	user, errUser := factories.UserFactory("get", &newtogo)

	if errUser != nil {
		context.JSON(http.StatusBadRequest, errUser.Error())
		return
	}

	if user != nil && user.CountTasks() >= user.LimitTasks {
		context.JSON(http.StatusBadRequest, "tasks has been limit.")
		return
	}

	result, errtogo := factories.TogoFactory("add", &newtogo)

	if errtogo != nil {
		context.JSON(http.StatusInternalServerError, errtogo.Error())
		return
	}

	context.JSON(http.StatusOK, result)
}
