package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kier1021/togo/api/controllers"
	"github.com/kier1021/togo/api/repositories"
	"github.com/kier1021/togo/api/services"
	"github.com/kier1021/togo/db"
)

type APIRoutes struct {
	dbs    *db.DB
	engine *gin.Engine
}

func NewAPIRoutes(dbs *db.DB) *APIRoutes {
	return &APIRoutes{
		dbs:    dbs,
		engine: gin.New(),
	}
}

func (routes *APIRoutes) GetEngine() *gin.Engine {
	return routes.engine
}

func (routes *APIRoutes) SetRoutes() {

	// Init user dependencies
	userRepo := repositories.NewUserRepository(routes.dbs.MongoDB)
	userSrv := services.NewUserService(userRepo)
	userCtrl := controllers.NewUserController(userSrv)

	// Init user task dependencies
	userTaskRepo := repositories.NewUserTaskRepository(routes.dbs.MongoDB)
	userTaskSrv := services.NewUserTaskService(userTaskRepo, userRepo)
	userTaskCtrl := controllers.NewUserTaskController(userTaskSrv)

	routes.engine.GET(
		"/",
		func() gin.HandlerFunc {
			return func(c *gin.Context) {
				c.JSON(200, map[string]interface{}{"message": "Hello World!!"})
			}
		}(),
	)

	routes.engine.POST(
		"/user",
		userCtrl.CreateUser(),
	)

	routes.engine.GET(
		"/users",
		userCtrl.GetUsers(),
	)

	routes.engine.PUT(
		"/user/task",
		userTaskCtrl.AddTaskToUser(),
	)

	routes.engine.GET(
		"/user/task",
		userTaskCtrl.GetTasksOfUser(),
	)
}
