package main

import (
	"pt.example/grcp-test/http/router"
)

func main() {
	r := router.New()
	r.Run() // listen and serve on 0.0.0.0:8080
}
