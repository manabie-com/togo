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
	userService := service.NewUserService(userRepo)
	userHttpController := httpInterface.NewUserController(userService)
	e := gin.New()
	userGroup := e.Group("/users")
	userGroup.POST("/register", userHttpController.Register)
	userGroup.POST("/login", userHttpController.Login)
	e.Run(fmt.Sprintf("0.0.0.0:%d", 8080))
}
