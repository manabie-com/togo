package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello\n")
	})

	http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil)
}
