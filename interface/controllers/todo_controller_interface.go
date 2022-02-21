package controllers

import (
	"github.com/gin-gonic/gin"
)

type TodoControllerInterface interface {
	GetAllTodoUser(c *gin.Context)
}
