package main

import (
	"os"

	lr "togo/utils/logger"

	"togo/config"
	"togo/controller"
	repository "togo/repository/task"
	"togo/router"
	"togo/service"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	connection     *mongo.Client             = config.ConnectMongo(os.Getenv("DATABASE_URI"), os.Getenv("DATABASE_PORT"))
	taskrepository repository.TaskRepository = repository.NewMongoRepository(connection)
	taskservice    service.TaskService       = service.NewTaskService(taskrepository)
	taskController controller.TaskController = controller.NewTaskController(taskservice)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {
	// Set logging
	logger := lr.NewLogger(os.Getenv("LOG_LEVEL"))
	port := os.Getenv("PORT")

	httpRouter.POST("/tasks", taskController.CreateTask)

	logger.Info().Msgf("Serving at %v", port)
	httpRouter.SERVE(port)
}
