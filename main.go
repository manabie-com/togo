package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/url"
	"togo/config"
	"togo/domain/service"
	"togo/infrastructure/persistent"
	"togo/interface/http"
	"togo/interface/http/middleware"
	configLib "togo/pkg/config"
)

func main() {
	var configFileName = flag.String("cfg", ".", "Configuration file path")
	flag.Parse()
	cfg := config.Config{}
	err := configLib.LoadConfig(*configFileName, "config.yaml", &cfg)
	if err != nil {
		log.Fatalln(err.Error())
	}
	mysqlCfg := cfg.MySQLConf
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", mysqlCfg.Username, url.QueryEscape(mysqlCfg.Password), mysqlCfg.Host, mysqlCfg.Port, mysqlCfg.AuthenticationDatabase))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {

		}
	}(db)
	userRepo := persistent.NewUserMysqlRepository(db)
	taskRepo := persistent.NewTaskMySQLRepository(db)
	tokenService := service.NewTokenService("sdasdasdasd")
	userService := service.NewUserService(userRepo, tokenService)
	taskSvc := service.NewTaskService(taskRepo)
	taskHttpController := httpInterface.NewTaskController(taskSvc)
	userHttpController := httpInterface.NewUserController(userService)
	e := gin.Default()
	userGroup := e.Group("/users")
	userGroup.POST("/register", userHttpController.Register)
	userGroup.POST("/login", userHttpController.Login)
	taskGroup := e.Group("/task").Use(middleware.AuthMiddleware(tokenService))
	taskGroup.POST("/create", taskHttpController.Create)
	e.Run(fmt.Sprintf("0.0.0.0:%d", 8080))
}
