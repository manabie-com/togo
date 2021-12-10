package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/manabie/project/config"
	"log"
	"time"
)

func main() {
	router := gin.Default()

	routerConfig := cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"X-Requested-With", "Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}

	router.Use(cors.New(routerConfig))

	conf, err := config.ReadConf("config/config-docker.yml")
	if err != nil {

	}


	if err := router.Run(conf.Server.PortServer); err != nil {
		log.Println("run port server failed port %s: ",conf.Server.PortServer)
	}
}