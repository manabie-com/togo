package main

import (
	"github.com/jmramos02/akaru/internal/api"
	"github.com/jmramos02/akaru/internal/database"
	"github.com/jmramos02/akaru/internal/model"
)

func main() {
	//initialize the database first
	db := database.Init()
	//run the migrations
	db.AutoMigrate(&model.User{}, &model.Task{})

	r := api.InitializeAPI()

	r.Run()
}
