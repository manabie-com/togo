package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shanenoi/togo/config"
	"github.com/shanenoi/togo/internal/storages/postgresql"
	"github.com/shanenoi/togo/internal/storages/redis"
	httpRoutes "github.com/shanenoi/togo/internal/transport/http"

	"net/http"
	"time"
)

func InitRouter(groupApis string, serverRoutes *gin.Engine) *gin.RouterGroup {
	serverRoutes.Use(gin.Logger())
	serverRoutes.Use(gin.Recovery())
	serverRoutes.Use(httpRoutes.CORSMiddleware())

	return serverRoutes.Group(groupApis)
}

func InitDatabase() *postgresql.PostgreSQLWrapper {
	databaseConfigs := config.DatabaseConfigs()
	return postgresql.Connect(databaseConfigs.DB_URI)
}

func InitRedis() *redis.RedisWrapper {
	redisConfigs := config.RedisConfigs()
	return redis.Connect(
		redisConfigs.REDIS_URI,
		redisConfigs.REDIS_PASS,
		redisConfigs.REDIS_DEFAULT_DB,
	)
}

func main() {
	serverConfigs := config.ServerConfigs()

	redisEngine := InitRedis()
	dbEngine := InitDatabase()
	postgresql.MakeMigrations(dbEngine.Database)
	defaultConfigs := &config.ThirdAppAdapter{
		Database: dbEngine,
		Redis:    redisEngine,
	}

	routeEngine := gin.New()
	routersInit := InitRouter(serverConfigs.GROUP_API, routeEngine)
	httpRoutes.ConfigTaskRouter(routersInit, defaultConfigs)
	httpRoutes.ConfigUserRouter(routersInit, defaultConfigs)

	server := &http.Server{
		Addr:           serverConfigs.PORT,
		Handler:        routeEngine,
		ReadTimeout:    serverConfigs.READTIMEOUT * time.Second,
		WriteTimeout:   serverConfigs.WRITETIMEOUT * time.Second,
		MaxHeaderBytes: serverConfigs.MAXHEADERBYTES,
	}

	server.ListenAndServe()
}
