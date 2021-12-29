package handler

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"togo/controller"
	"togo/db"
	"togo/form"
	"togo/middleware"
	"togo/model"
)

type APIEnv struct {
	DB *gorm.DB
}

func (a *APIEnv) GetTask(c *gin.Context) {
	createdDate := c.Request.URL.Query()["created_date"][0]
	user := middleware.GetUserFromCtx(c)

	listTask, err := controller.GetAllTaskByUser(user.Id, createdDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": listTask})
}

func (a *APIEnv) CreateTask(c *gin.Context) {
	var input form.Task
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := middleware.GetUserFromCtx(c)

	task, err := controller.AddTask(user, input.Content)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.CreateTask(task)

	c.JSON(http.StatusOK, gin.H{"data": task})
}

func (a *APIEnv) Register(c *gin.Context) {
	var input form.User
	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if input.MaxTodo == 0 {
		input.MaxTodo = 5
	}

	existedUser := model.User{}
	db.DB.First(&existedUser, "username = ?", input.Username)

	if existedUser.Id != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user existed"})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
		MaxTodo:  input.MaxTodo,
	}
     controller.CreateUser(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}
