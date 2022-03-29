package controllers

import (
	"errors"
	"net/http"
	"togo/globals/validator"
	"togo/models"
	"togo/models/form"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Index(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message":"To-do API"})
}

func CreateTodo(ctx *gin.Context) {
	form := form.Form{}
	ctx.ShouldBind(&form)
	errJSONS := validator.Validate(form)
	if errJSONS != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message":errJSONS})
		return
	}

	user, err := models.UserById(form.UserID)
	if err != nil {
		if (errors.Is(err, gorm.ErrRecordNotFound)) {
			user, err = models.CreateUserWithTask(form.UserID, form.TaskDetail)
			if user.ID != 0 {
				ctx.JSON(http.StatusOK, gin.H{
					"message": "Create ToDo successful",
					"task": user.Tasks[0],
				})
				return
			}
		}
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, err.Error())
			return
		}
	}

	if max, err := models.UserAtDailyLimit(form.UserID); max || err != nil {
		if max {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Daily limit reached"})
		} else {
			ctx.JSON(http.StatusInternalServerError, err.Error())
		}
		return
	}

	task := models.Task{ UserID: user.ID, Detail: form.TaskDetail}
	createdTask, err := models.CreateTask(task)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Create ToDo successful",
		"task": createdTask,
	})
}