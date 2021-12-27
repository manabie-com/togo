package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"togo/db"
	"togo/form"
	"togo/middleware"
	"togo/model"
)

func ListTask(c *gin.Context)  {
	createdDate := c.Request.URL.Query()["created_date"][0]
	user := middleware.GetUserFromCtx(c)

	listTask,err :=GetAllTaskByUser(user.Id,createdDate)

	if err != nil {
		c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": listTask})
}

func GetAllTaskByUser(userId int,createdDate string) ([]*model.Task,error) {
	var task []*model.Task

	if err := db.DB.Where("user_id = ? AND created_date =?",strconv.Itoa(userId),createdDate).Find(&task).Error; err != nil {
		return nil,fmt.Errorf("Record not found")
	}

	return task,nil
}

func CreateTask(c *gin.Context) {
	var input form.Task
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserFromCtx(c)

	task,err := addTask(user,input.Content)

	if err != nil {
		 c.JSON(http.StatusCreated, gin.H{"error": err.Error()})
		 return
	}

	db.DB.Create(&task)

	c.JSON(http.StatusCreated, gin.H{"data": task})
}

func addTask(user model.User,content string) (*model.Task,error) {
	current := time.Now()

	task := &model.Task{
		Content:     content,
		UserID:      user.Id,
		CreatedDate: current.Format("2006-01-02"),
	}

	listTask,err := GetAllTaskByUser(user.Id,task.CreatedDate)
	if err != nil {
		return nil,err
	}

	if len(listTask) >=user.MaxTodo {
		return nil,fmt.Errorf("user reach task limit error")
	}

	return task,nil
}
