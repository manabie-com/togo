package main

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/configs"
	_ "github.com/manabie-com/togo/internal/docs"
	"github.com/manabie-com/togo/internal/providers"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"os"
)

func main() {
	// Read App Config
	appConfig := configs.NewApplicationConfig()

	// Init Logger
	logger := logrus.New()
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.InfoLevel)

	// Init DB Connection
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})
	if err != nil {
		logger.Errorf("Connect to DB error: %s", err.Error())
		panic(err)
	}

	// Init API Handlers
	appMiddleware := providers.ProvideApplicationMiddleware(appConfig.Jwt.SecretKey)
	userHandler := providers.ProvideUserHandler(db, appConfig.Jwt.SecretKey)
	taskHandler := providers.ProvideTaskHandler(db)
	configurationHandler := providers.ProvideConfigurationHandler(db)

	// Init Gin Server
	app := gin.Default()
	app.Use(appMiddleware.ApplyCorsFilter()).
		Use(appMiddleware.ApplyJwtFilter())

	app.GET("/login", userHandler.GetAuthToken)
	taskRouter := app.Group("/tasks")
	{
		taskRouter.POST("/", taskHandler.CreateTask)
		taskRouter.GET("/", taskHandler.GetListTask)
	}

	configurationRouter := app.Group("/configurations")
	{
		configurationRouter.POST("/", configurationHandler.CreateConfiguration)
		configurationRouter.GET("/date", configurationHandler.GetConfigurationByDate)
	}
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	if err := app.Run(appConfig.Server.Address); err != nil {
		logger.Errorf("Start HTTP Server error: %s", err.Error())
		panic(err)
	}
}
