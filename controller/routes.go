package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"log"
	"manabie-com/togo/controller/healthy"
	"manabie-com/togo/controller/user"
	"manabie-com/togo/global"
	_ "manabie-com/togo/swagger"
	"net/http"
)

func Initialize() {
	var port = global.Config.ServerPort

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.MaxMultipartMemory = 8 << 20 // 8 MiB

	healthy.RegisterRoutes(router)
	user.RegisterRoutes(router)
	// enable swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("/swagger/doc.json")))

	fmt.Println("Interview backend API listening on port: " + port)
	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
