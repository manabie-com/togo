package router

import (
	"encoding/json"
	"net/http"

	"github.com/manabie-com/togo/internal/services/transport"
)

func (route *Router) getAuthToken(resp http.ResponseWriter, req *http.Request) {

	var (
		method = "get-auth"
		err    error
	)

	data, err := transport.Init(route.Conn, method, req)
	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": data["token"],
	})
}

func (route *Router) listTasks(resp http.ResponseWriter, req *http.Request) {

	var (
		method = "list-tasks"
		err    error
	)

	data, err := transport.Init(route.Conn, method, req)
	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": data["tasks"],
	})
}

func (route *Router) addTask(resp http.ResponseWriter, req *http.Request) {

	var (
		method = "add-task"
		err    error
	)

	data, err := transport.Init(route.Conn, method, req)
	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(data["status"].(int))
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": data["task"],
	})
}
