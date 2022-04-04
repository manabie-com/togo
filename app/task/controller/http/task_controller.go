package http

import (
	"github.com/ansidev/togo/auth/constant"
	authDto "github.com/ansidev/togo/auth/dto"
	authMiddleware "github.com/ansidev/togo/auth/middleware"
	authService "github.com/ansidev/togo/auth/service"
	"github.com/ansidev/togo/task/dto"
	"github.com/ansidev/togo/task/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func registerRoutes(router *gin.Engine, taskController *taskController, authService authService.IAuthService) {
	v1 := router.Group("/task/v1")

	v1.POST("/tasks", authMiddleware.AuthUser(authService), taskController.CreateTask)
}

func NewTaskController(
	router *gin.Engine,
	authService authService.IAuthService,
	taskService service.ITaskService,
) {
	controller := &taskController{authService, taskService}
	registerRoutes(router, controller, authService)
}

type taskController struct {
	authService authService.IAuthService
	taskService service.ITaskService
}

func (ctrl *taskController) CreateTask(ctx *gin.Context) {
	var request dto.CreateTaskRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.Error(err)
		return
	}

	ac, _ := ctx.Get(constant.AuthKey)
	authenticationCredential := ac.(authDto.UserCredential)

	task, err := ctrl.taskService.Create(request, authenticationCredential.ID)

	if err != nil {
		ctx.Error(err)
		return
	}

	ctx.JSON(http.StatusCreated, task)
}
