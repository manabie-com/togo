package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jfzam/togo/infrastructure/auth"
	"github.com/jfzam/togo/infrastructure/persistence"
	"github.com/jfzam/togo/interfaces"
	"github.com/jfzam/togo/interfaces/middleware"
	"github.com/joho/godotenv"
)

func init() {
	//To load our environmental variables.
	if err := godotenv.Load(); err != nil {
		log.Println("no env gotten")
	}
}

func main() {

	//dbdriver := os.Getenv("DB_DRIVER")
	host := os.Getenv("DB_HOST")
	password := os.Getenv("DB_PASSWORD")
	user := os.Getenv("DB_USER")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	//redis details
	redis_host := os.Getenv("REDIS_HOST")
	redis_port := os.Getenv("REDIS_PORT")
	redis_password := os.Getenv("REDIS_PASSWORD")

	services, err := persistence.NewRepositories(user, password, port, host, dbname)
	if err != nil {
		panic(err)
	}
	defer services.Close()
	services.Automigrate()

	redisService, err := auth.NewRedisDB(redis_host, redis_port, redis_password)
	if err != nil {
		log.Fatal(err)
	}

	tk := auth.NewToken()

	users := interfaces.NewUsers(services.User, redisService.Auth, tk)
	Tasks := interfaces.NewTask(services.Task, services.User, redisService.Auth, tk)
	authenticate := interfaces.NewAuthenticate(services.User, redisService.Auth, tk)

	r := gin.Default()
	r.Use(middleware.CORSMiddleware()) //For CORS

	//user routes
	r.POST("/users", users.SaveUser)
	r.GET("/users", users.GetUsers)
	r.GET("/users/:user_id", users.GetUser)

	//post routes
	r.POST("/Task", middleware.AuthMiddleware(), middleware.MaxSizeAllowed(8192000), Tasks.SaveTask)
	r.GET("/Task/:Task_id", Tasks.GetTaskAndCreator)
	r.GET("/Task", Tasks.GetAllTask)

	//authentication routes
	r.POST("/login", authenticate.Login)
	r.POST("/logout", authenticate.Logout)
	r.POST("/refresh", authenticate.Refresh)

	//Starting the application
	app_port := os.Getenv("PORT") //using heroku host
	if app_port == "" {
		app_port = "8888" //localhost
	}
	log.Fatal(r.Run(":" + app_port))
}
