package main

import (
	"github.com/gin-gonic/gin"
	"todo/src/controller"
)

/*
Router will use some controllers based on the objects
 */

type Router struct {
	Con *controller.Controller
	Gin *gin.Engine
}

func NewRouters() *Router {
	return &Router{
		Gin: gin.Default(),
		Con: controller.NewController(),
	}
}

func (r *Router) ConnectController() {
	// controller for user
	r.Con.UserController(r.Gin)
}

