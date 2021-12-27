package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"togo/db"
	"togo/form"
	"togo/model"
)

func Register(c *gin.Context) {
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
	db.DB.Create(&user)

	c.JSON(http.StatusCreated, gin.H{"data": user})
}
