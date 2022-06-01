package main

import (
	"os"

	lr "togo/utils/logger"

	"togo/config"
	"togo/controller"
	taskRepo "togo/repository/task"
	userRepo "togo/repository/user"
	"togo/router"
	"togo/service"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	connection     *mongo.Client             = config.ConnectMongo(os.Getenv("DATABASE_URI"), os.Getenv("DATABASE_PORT"))
	taskrepository taskRepo.TaskRepository   = taskRepo.NewMongoRepository(connection)
	userrepository userRepo.UserRepository   = userRepo.NewMongoRepository(connection)
	taskservice    service.TaskService       = service.NewTaskService(taskrepository)
	userservice    service.UserService       = service.NewUserService(userrepository)
	taskController controller.TaskController = controller.NewTaskController(taskservice)
	userController controller.UserController = controller.NewUserController(userservice)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	// Set logging
	logger := lr.NewLogger(os.Getenv("LOG_LEVEL"))
	port := os.Getenv("PORT")

	httpRouter.POST("/tasks", taskController.CreateTask)
	httpRouter.POST("/register", userController.Register)
	httpRouter.PUT("/login", userController.Login)

	logger.Info().Msgf("Serving at %v", port)
	httpRouter.SERVE(port)
}
