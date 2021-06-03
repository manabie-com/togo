package api

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// Swagger
	_ "github.com/swaggo/gin-swagger/example/basic/docs"
)

// Route ...
type Route struct {
	method  string
	path    string
	handler gin.HandlerFunc
}

// AuthenRoute ...
type AuthenRoute struct {
	method  string
	path    string
	authen  gin.HandlerFunc
	handler gin.HandlerFunc
}

var listRoute = []Route{
	{http.MethodGet, "/login", LoginUser},
}

var listAuthenRoute = []AuthenRoute{
	{http.MethodGet, "/tasks", authenMiddleware, GetListTasks},
	{http.MethodPost, "/tasks", authenMiddleware, AddTask},
}

// NewRouter ...
func NewRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	for _, r := range listRoute {
		router.Handle(r.method, r.path, r.handler)
	}

	for _, r := range listAuthenRoute {
		router.Handle(r.method, r.path, r.authen, r.handler)
	}
	return router
}
