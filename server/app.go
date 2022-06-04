package main

import (
	"fmt"
	"net/http"
	"os"

	lr "togo/utils/logger"

	"togo/config"
	"togo/controller"
	taskRepo "togo/repository/task"
	userRepo "togo/repository/user"
	"togo/router"
	"togo/service"

	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	logger         zerolog.Logger            = lr.NewLogger(os.Getenv("LOG_LEVEL"))
	connection     *mongo.Client             = config.ConnectMongo(os.Getenv("DATABASE_URI"), os.Getenv("DATABASE_PORT"))
	taskrepository taskRepo.TaskRepository   = taskRepo.NewMongoRepository(connection)
	userrepository userRepo.UserRepository   = userRepo.NewMongoRepository(connection)
	taskservice    service.TaskService       = service.NewTaskService(taskrepository, userrepository)
	userservice    service.UserService       = service.NewUserService(userrepository)
	taskController controller.TaskController = controller.NewTaskController(taskservice, logger)
	userController controller.UserController = controller.NewUserController(userservice, logger)
	httpRouter     router.RouterInterface    = router.NewChiRouter(taskController, userController)
)

func main() {
	// Set logging
	port := fmt.Sprintf(":%v", os.Getenv("PORT"))

	logger.Info().Msgf("Serving at %v", port)
	http.ListenAndServe(port, httpRouter.Router())
}
