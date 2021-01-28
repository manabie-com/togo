package services

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/usecase"
	"github.com/manabie-com/togo/utils"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey      string
	TaskUsecase *usecase.TaskUsecase
	UserUsecase *usecase.UserUsecase
}

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
		req, ok = utils.ValidToken(req, s.JWTKey)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			s.listTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		case http.MethodDelete:
			s.deleteTasks(resp, req)
		}
		return
	}
}

func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := utils.UserIDFromCtx(req.Context())
	tasks, err := s.TaskUsecase.ListTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		utils.Value(req, "created_date"),
	)
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
	return
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := utils.Value(req, "user_id")

	if !s.UserUsecase.ValidateUser(req.Context(), id, utils.Value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := utils.CreateToken(id.String, s.JWTKey)
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

func (s *ToDoService) deleteTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := utils.UserIDFromCtx(req.Context())

	err := s.TaskUsecase.DeleteTasks(req.Context(), sql.NullString{
		String: id,
		Valid:  true,
	})

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]bool{
		"success": true,
	})
	return
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := utils.UserIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	// Check max tasks per day
	isMaxTask, err := s.TaskUsecase.CheckMaxTaskPerDay(req, userID, t.CreatedDate)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	if isMaxTask {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "limited to create task per day",
		})
		return
	}

	resp.Header().Set("Content-Type", "application/json")

	err = s.TaskUsecase.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": t,
	})
	return
}
