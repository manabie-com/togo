package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kier1021/togo/api/controllers"
	"github.com/kier1021/togo/api/services"
)

type APIRoutes struct {
	engine *gin.Engine
}

func NewAPIRoutes() *APIRoutes {
	return &APIRoutes{
		engine: gin.New(),
	}
}

func (routes *APIRoutes) GetEngine() *gin.Engine {
	return routes.engine
}

func (routes *APIRoutes) SetRoutes() {

	taskSrv := services.NewTaskService()
	taskCtrl := controllers.NewTaskController(taskSrv)

	routes.engine.GET(
		"/",
		func() gin.HandlerFunc {
			return func(c *gin.Context) {
				c.JSON(200, map[string]interface{}{"message": "Hello World!!"})
			}
		}(),
	)

	routes.engine.POST(
		"/tasks",
		taskCtrl.CreateTasks(),
	)
}
