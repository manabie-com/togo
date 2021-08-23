package app

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"gorm.io/gorm"
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/internal/controllers"
	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/config"
	"github.com/manabie-com/togo/internal/db"
	"github.com/manabie-com/togo/internal/middleware"
)

type APP struct {
	Router *gin.Engine
	Server *http.Server
}

func InitTaskController(db *gorm.DB) *controllers.TaskController {
	repository := repositories.ProvideTaskRepository(db)
	service := services.ProvideTaskService(repository)

	return &controllers.TaskController{
		TaskService: service,
	}
}

func InitAuthController(db *gorm.DB) *controllers.AuthController {
	repository := repositories.ProvideUserRepository(db)
	service := services.ProvideAuthService(repository)

	return &controllers.AuthController{
		AuthService: service,
	}
}

// Initialize : Initialize Application Components
func (app *APP) Initialize() {
	gin.SetMode(gin.ReleaseMode)

	config.InitConfig()

	db.InitDB(
		config.Config.DBHost,
		config.Config.DBPort,
		config.Config.DBUser,
		config.Config.DBPass,
		config.Config.DBName,
	)

	app.Router = gin.Default()
	app.Router.RedirectTrailingSlash = true
	app.Router.RedirectFixedPath = true
	app.Router.Use(middleware.EnableCORS())
	app.Router.Use(middleware.ErrorHandler)

	v1 := app.Router.Group("/v1")
	{
		taskRoutes := v1.Group("/tasks")
		taskCtrl := InitTaskController(db.DB)
		taskRoutes.GET("", middleware.AuthUser(), taskCtrl.FindAll)
		taskRoutes.POST("", middleware.AuthUser(), taskCtrl.Create)
		// taskRoutes.GET("/:id", taskCtrl.FindByID)
		// taskRoutes.PUT("/:id", taskCtrl.Update)
		// taskRoutes.DELETE("/:id", taskCtrl.Delete)
			
		authRoutes := v1.Group("/auth")
		authCtrl := InitAuthController(db.DB)
		authRoutes.POST("/signup", authCtrl.Signup)
		authRoutes.POST("/login", authCtrl.Login)
	}

	app.Server = &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(config.Config.AppPort),
		Handler: app.Router,
	}
}

// Serve : Serve the Application with Error Channels
func (app *APP) Serve() {
	errChan := make(chan error, 10)

	go func() {
		errChan <- app.Server.ListenAndServe()
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case err := <-errChan:
			if err != nil {
				log.Fatal(err)
			}
		case s := <-signalChan:
			log.Println(fmt.Sprintf("Captured message %v. Exiting...", s))
			os.Exit(0)
		}
	}
}
