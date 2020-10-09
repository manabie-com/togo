package module

import (
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/module/auth"
	"github.com/manabie-com/togo/internal/module/task"
	"github.com/manabie-com/togo/internal/module/user"
	"gorm.io/gorm"
)

// LoadModules func
func LoadModules(e *echo.Echo, db *gorm.DB) {
	//TODO: enhance init controller by go.uber.org/fx

	userRepo, _ := user.NewUserRepository(db)
	userService, _ := user.NewUserService(userRepo)

	authController, _ := auth.NewAuthController(userService)
	auth.LoadRoute(e, authController)

	userController, _ := user.NewUserController(userService)
	user.LoadRoute(e, userController)

	taskRepo, _ := task.NewTaskRepository(db)
	taskService, _ := task.NewTaskService(taskRepo)
	taskController, _ := task.NewTaskController(taskService, userService)
	task.LoadRoute(e, taskController)
}
