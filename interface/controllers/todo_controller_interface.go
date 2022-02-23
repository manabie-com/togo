package controllers

import (
	"github.com/gin-gonic/gin"
)

type TodoControllerInterface interface {
	CreateTodoUser(c *gin.Context)
	GetAllTodoUser(c *gin.Context)
}
