package main

import (
	"github.com/gin-gonic/gin"

	"github.com/manabie-com/togo/controllers"
)

func main() {
	hostDefault()
}

func hostDefault() {

	r := gin.Default()

	r.POST("/todo/add", controllers.AddTodoTask)

	r.Run()
}
