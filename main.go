package main

import (
	"context"

	"github.com/joho/godotenv"
	"pt.example/grcp-test/database"
	"pt.example/grcp-test/http/router"
)

func main() {
	var err error

	err = godotenv.Load()

	if err != nil {
		panic(err)

	}
	err = database.Init()

	if err != nil {
		panic(err)
	}

	defer func() {
		if err := database.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := router.New()
	r.Run() // listen and serve on 0.0.0.0:8080
}
