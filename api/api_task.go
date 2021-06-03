package api

import (
	"net/http"
	"time"
	"togo/model"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddTask(c *gin.Context) {
	var task model.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	task.ID = uuid.New().String()
	task.UserID = c.Request.Header.Get("user_id")
	task.CreatedDate = time.Now().Format("2006-01-02")
	data, err := model.AddTask(task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func GetListTasks(c *gin.Context) {
	var task model.Task
	err := c.BindJSON(&task)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	task.UserID = c.Request.Header.Get("user_id")
	data, err := model.GetListTasks(task)
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
