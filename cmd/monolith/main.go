package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
	"togo/internal/pkg/db"
	"togo/internal/services/task/api"
	"togo/internal/services/task/application"
	"togo/internal/services/task/store/postgres"
)

func main() {
	logrus.Info("Starting monolith")

	if _, err := os.Stat(".env"); err == nil {
		err = godotenv.Load()
		if err != nil {
			logrus.Fatal("Cannot load env: ", err)
		}
	}

	postgresDB := db.DB{}
	postgresDB.Connect()

	router := createMonolith(&postgresDB)
	router.GET("/healthz", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	server := &http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: router,
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logrus.Println("receive interrupt signal")
		if err := server.Close(); err != nil {
			logrus.Fatal("Server Close:", err)
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			logrus.Println("Server closed under request")
		} else {
			logrus.Fatal("Server closed unexpect")
		}
	}

	logrus.Println("Server exiting")
}
func createMonolith(db *db.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"*"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{
		"Access-Control-Allow-Origin",
		"Origin",
		"Content-Type",
		"Content-Length",
		"Accept-Encoding",
		"X-CSRF-Token",
		"Authorization",
		"Referer",
		"X-Size",
	}
	r.Use(cors.New(corsConfig))

	userRepo := postgres.NewUserRepository(db)
	userService := application.NewUserService(userRepo)

	taskRepo := postgres.NewTaskRepository(db)
	taskService := application.NewTaskService(taskRepo, userRepo)

	userAPI := api.NewUserAPI(userService)
	taskAPI := api.NewTaskAPI(taskService)
	userAPI.AddRoutes(r)
	taskAPI.AddRoutes(r)
	return r
}
