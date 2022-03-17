package server

import (
	"github.com/gin-gonic/gin"
	"github.com/kier1021/togo/api/controllers"
	"github.com/kier1021/togo/api/repositories"
	"github.com/kier1021/togo/api/services"
	"github.com/kier1021/togo/db"
)

// APIRoutes holds the routes of API
type APIRoutes struct {
	dbs    *db.DB
	engine *gin.Engine
}

// NewAPIRoutes is the constructor for APIRoutes
func NewAPIRoutes(dbs *db.DB) *APIRoutes {
	return &APIRoutes{
		dbs:    dbs,
		engine: gin.New(),
	}
}

// GetEngine returns the gin engine
func (routes *APIRoutes) GetEngine() *gin.Engine {
	return routes.engine
}

// SetRoutes set the endpoints of API
func (routes *APIRoutes) SetRoutes() {

	// Init user dependencies
	userRepo := repositories.NewUserRepository(routes.dbs.MongoDB)
	userSrv := services.NewUserService(userRepo)
	userCtrl := controllers.NewUserController(userSrv)

	// Init user task dependencies
	userTaskRepo := repositories.NewUserTaskRepository(routes.dbs.MongoDB)
	userTaskSrv := services.NewUserTaskService(userTaskRepo, userRepo)
	userTaskCtrl := controllers.NewUserTaskController(userTaskSrv)

	// Set the endpoints
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
