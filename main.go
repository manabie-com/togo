package main

import (
	"fmt"

	"github.com/manabie-com/togo/configs"
	"github.com/manabie-com/togo/controllers"
	"github.com/manabie-com/togo/repositories"
	"github.com/manabie-com/togo/server"
	"github.com/manabie-com/togo/services"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                     = configs.SetupDatabaseConnection()
	taskRepository repositories.ITaskRepository = repositories.NewTaskRepository(db)
	userRepository repositories.IUserRepository = repositories.NewUserRepository(db)
	taskService    services.ITaskService        = services.NewTaskService(taskRepository, userRepository)
	taskController controllers.ITaskController  = controllers.NewTaskController(taskService)
)

func main() {
	server, err := server.NewServer(taskController)

	if err != nil {
		fmt.Println("ERROR occcurred when create server", err)
	}

	errWhenStart := server.StartServer("localhost:12345")

	if errWhenStart != nil {
		fmt.Println("ERROR occcurred when start server", errWhenStart)
	}
}
