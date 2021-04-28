package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/validate"
)

// ToDoService implement HTTP server
type Router struct {
	JWTKey string
	Conn   *gorm.DB
}

func (route *Router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		{
			route.getAuthToken(resp, req)
			return
		}
	case "/tasks":
		{
			req, err := validate.ValidToken(req)

			if err != nil {

				resp.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(resp).Encode(map[string]string{
					"error": err.Error(),
				})

				return

			}

			switch req.Method {
			case http.MethodGet:

				route.listTasks(resp, req)

			case http.MethodPost:

				route.addTask(resp, req)

			}

			return
		}
	}
}
