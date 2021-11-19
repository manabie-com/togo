package main

import (
	"manabie/manabie/middlewares"
	"manabie/manabie/routers"
	"os"

	"github.com/gin-contrib/location"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

func getPort() string {
	p := os.Getenv("HOST_PORT")
	if p != "" {
		return ":" + p
	}
	return ":8888"
}

func main() {
	port := getPort()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := gin.Default()
	r.Use(location.Default())
	r.Use(middlewares.CORSMiddleware())
	//rg := r.Group("manabie/v1")
	rg := r.Group("manabie")
	rg.Use(middlewares.CORSMiddleware())
	{
		routers.UserAPIRoute(rg)
	}
	rg.Use(middlewares.CORSMiddleware(), middlewares.ValidateToken())
	{
		routers.TodoAPIRoute(rg)
	}
	r.Run(port)
}
