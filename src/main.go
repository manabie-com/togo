package main

import (
	"fmt"
	"net/http"
)

func main() {
	a := App{
		volumeServer: "http://localhost:3001",
		lock:         make(map[string]struct{}),
		volume:       "/tmp/volume1/",
	}
	fmt.Println("Starting master server")
	http.ListenAndServe(fmt.Sprintf(":%d", 3000), &a)
}
