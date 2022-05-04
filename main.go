package main

import (
	"github.com/gin-gonic/gin"

	"github.com/manabie-com/togo/controllers"
	"github.com/manabie-com/togo/database"
)

func main() {
	hostDefault()
}

func hostDefault() {

	r := gin.Default()

	database.ConnectDatabase()

	r.POST("/todo/add", controllers.AddTogoTask)

	r.Run()
}
