package main

import (
	"fmt"
	"net/http"

	"togo/database"

	"togo/src/modules/auth"
	"togo/src/modules/tasks"
	"togo/src/modules/users"

	"github.com/gin-gonic/gin"

	JwtMiddleware "togo/src/middleware/jwt"
)

func setupRouter() *gin.Engine {
	// Creates a router without any middleware by default
	app := gin.New()

	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	app.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	app.Use(gin.Recovery())

	// Ping test
	app.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// initializes database
	db, _ := database.Initialize()
	app.Use(database.Inject(db))
	app.Use(JwtMiddleware.ParseUserContext())
	users.RegisterRouter(app)
	auth.RegisterRouter(app)
	tasks.RegisterRouter(app)
	// users := router.Group("/users")
	// {
	// 	users.POST("/", usersController.Create)
	// }
	// router.POST("/", controller.Create)

	return app
}

func main() {
	server := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	fmt.Println("Listen and Server in 0.0.0.0:8080")
	server.Run(":8080")
}
