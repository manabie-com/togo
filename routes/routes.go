package routes

import (
	"fmt"
	"net/http"
	"strings"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	args := strings.Split(r.URL.Path, "/")
	fmt.Println(len(args), args)

	switch args[1] {
	case "user":
		UserRoute(w, r)
	case "task":
		TaskRoute(w, r)
	case "auth":
		AuthRoute(w, r)
	default:
		w.Write([]byte("lntvan166: Hello"))
	}
}
