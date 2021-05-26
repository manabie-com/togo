package services

import (
	"encoding/json"
	"github.com/manabie-com/togo/internal/models"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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
		s.getAuthToken(resp, req)
		return
	case "/tasks":
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			s.listTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		}
		return
	}
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := req.FormValue("user_id")
	password := req.FormValue("password")

	if !s.Store.ValidateUser(req.Context(), id, password) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})

		return
	}

	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
		"data": token,
	})
}

func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		id,
		req.FormValue("created_date"),
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	json.NewEncoder(resp).Encode(map[string][]*models.Task{
		"data": tasks,
	})
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	t := &models.Task{}
	err := json.NewDecoder(req.Body).Decode(t)

	defer req.Body.Close()

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	user, getUserErr := s.Store.GetUser(req.Context(), id)

	if getUserErr != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": getUserErr.Error(),
		})

		return
	}

	tasks, getListTaskErr := s.Store.RetrieveTasks(
		req.Context(),
		id,
		t.CreatedDate,
	)

	if getListTaskErr != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": getListTaskErr.Error(),
		})

		return
	}

	if len(tasks) >= user.MaxTodo {
		resp.WriteHeader(http.StatusForbidden)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "User has reached today's limit of tasks.",
		})

		return
	}

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})

		return
	}

	json.NewEncoder(resp).Encode(map[string]*models.Task{
		"data": t,
	})
}
