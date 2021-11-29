package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"togo/config"
	"togo/models"
	"togo/router"
)

func main()  {
	cf := config.InitFromFile("")
	closeFunc, errFile := models.InitFromSQLLite(cf.DbConnection)
	if errFile != nil {
		log.Fatalf("Read file Database error!")
	}

	routersInit := router.InitRouter(cf.EnvironmentPrefix)
	maxHeaderBytes := 1 << 20
	endPoint := fmt.Sprintf(":%d", cf.ServerPort)
	server := &http.Server{
		Addr: endPoint,
		Handler: routersInit,
		ReadTimeout: time.Minute,
		WriteTimeout: time.Minute,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] Start http server listening %s", endPoint)
	_ = server.ListenAndServe()
	defer closeFunc()
}