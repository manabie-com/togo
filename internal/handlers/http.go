package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/manabie-com/togo/internal/services"
	"github.com/manabie-com/togo/internal/storages"
)

type HttpHandler struct {
	UserService *services.UserService
	TaskService *services.TaskService
}

func (h *HttpHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		token, err := h.UserService.GetAuthToken(req.Context(), req.FormValue("user_id"), req.FormValue("password"))
		if err != nil {
			resp.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}

		json.NewEncoder(resp).Encode(map[string]string{
			"data": token,
		})

		return
	case "/tasks":
		id, err := h.UserService.ValidToken(req.Header.Get("Authorization"))

		if err != nil {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		req = req.WithContext(context.WithValue(req.Context(), services.UserAuthKey(0), id))
		userID, _ := h.UserService.UserIDFromCtx(req.Context())

		switch req.Method {
		case http.MethodGet:
			tasks, err := h.TaskService.ListTasks(req.Context(), userID, req.FormValue("created_date"))

			resp.Header().Set("Content-Type", "application/json")

			if err != nil {
				resp.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(resp).Encode(map[string]string{
					"error": err.Error(),
				})
				return
			}

			json.NewEncoder(resp).Encode(map[string][]*storages.Task{
				"data": tasks,
			})
		case http.MethodPost:
			defer req.Body.Close()

			resp.Header().Set("Content-Type", "application/json")

			err := h.TaskService.IsReachedLimit(req.Context(), userID)
			if err != nil {
				resp.WriteHeader(err.Code)
				json.NewEncoder(resp).Encode(map[string]string{
					"error": err.Error(),
				})
				return
			}

			task, err := h.TaskService.AddTask(req.Context(), req.Body, userID)
			if err != nil {
				resp.WriteHeader(err.Code)
				json.NewEncoder(resp).Encode(map[string]string{
					"error": err.Error(),
				})
			}

			json.NewEncoder(resp).Encode(map[string]*storages.Task{
				"data": task,
			})
		}

		return
	}
}
