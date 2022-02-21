package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
	"togo/infrastructure/database"
	"togo/infrastructure/routes"
	"togo/registry"
)

var Mode = os.Getenv("MODE")
var port = os.Getenv("PORT")
var dbHost = os.Getenv("DB_HOST")
var dbPort = os.Getenv("DB_PORT")
var dbDatabase = os.Getenv("DB_DATABASE")
var dbUsername = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")

var db *database.DbGormStruct

func init() {
	gin.SetMode(Mode)
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)
	db = database.Init(dbInfo)
}

func main() {
	router := gin.Default()
	reg := registry.NewRegistry(db)
	router = routes.NewRouter(router, reg.NewAppController())

	router.Run(":" + port)
}
