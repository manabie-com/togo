package main

import (
	"context"

	"pt.example/grcp-test/database"
	"pt.example/grcp-test/http/router"
)

func main() {
	database.Init()

	defer func() {
		if err := database.Client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	r := router.New()
	r.Run() // listen and serve on 0.0.0.0:8080
}
