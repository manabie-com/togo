package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/manabie-com/togo/internal/storages"
)

// ToDoService implement HTTP server
type ToDoService struct {
	Router *mux.Router
	JWTKey string
	Store  storages.IToGoDB
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

	s.Router.ServeHTTP(resp, req)
}

//LoginHandler ...
func (s *ToDoService) LoginHandler(resp http.ResponseWriter, req *http.Request) {
	s.getAuthToken(resp, req)
}

// GetTasksHandler ...
func (s *ToDoService) GetTasksHandler(resp http.ResponseWriter, req *http.Request) {
	s.listTasks(resp, req)
}

// CreateTaskHandler ...
func (s *ToDoService) CreateTaskHandler(resp http.ResponseWriter, req *http.Request) {
	if !s.canAddTask(resp, req) {
		resp.WriteHeader(http.StatusNotAcceptable)
		return
	}
	s.addTask(resp, req)
}

// UpdateTaskStatusHandler ...
func (s *ToDoService) UpdateTaskStatusHandler(resp http.ResponseWriter, req *http.Request) {
	s.updateTaskStatus(resp, req)
}

// UpdateAllTaskStatusHandler ...
func (s *ToDoService) UpdateAllTaskStatusHandler(resp http.ResponseWriter, req *http.Request) {
	s.updateAllTasksStatus(resp, req)
}

// DeleteTaskHandler ...
func (s *ToDoService) DeleteTaskHandler(resp http.ResponseWriter, req *http.Request) {
	s.deleteTask(resp, req)
}

// DeleteTasksHandler ...
func (s *ToDoService) DeleteTasksHandler(resp http.ResponseWriter, req *http.Request) {
	s.deleteTasks(resp, req)
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	if !s.Store.ValidateUser(req.Context(), id, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id.String)
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
		sql.NullString{
			String: id,
			Valid:  true,
		},
		value(req, "created_date"),
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
}

func (s *ToDoService) canAddTask(resp http.ResponseWriter, req *http.Request) bool {
	userID, _ := userIDFromCtx(req.Context())
	maxTask, err := s.Store.GetUserMaxTask(
		req.Context(),
		sql.NullString{
			String: userID,
			Valid:  true,
		})
	if err != nil {
		log.Println(err)
		return false
	}

	countTodayTask, err := s.Store.GetUserTodayTask(
		req.Context(),
		sql.NullString{
			String: userID,
			Valid:  true,
		},
	)

	if err != nil {
		log.Println(err)
		return false
	}

	if countTodayTask < maxTask {
		return true
	}
	return false
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
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)
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
}

func (s *ToDoService) updateTaskStatus(resp http.ResponseWriter, req *http.Request) {
	userID, _ := userIDFromCtx(req.Context())
	vars := mux.Vars(req)
	taskID := vars["id"]
	status := value(req, "status")
	err := s.Store.UpdateStatusTask(
		req.Context(),
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: taskID,
			Valid:  true,
		},
		status)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
}

func (s *ToDoService) updateAllTasksStatus(resp http.ResponseWriter, req *http.Request) {
	userID, _ := userIDFromCtx(req.Context())
	createdDate := value(req, "created_date")
	status := value(req, "status")
	err := s.Store.UpdateAllStatusTasks(
		req.Context(),
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		createdDate,
		status)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
}

func (s *ToDoService) deleteTask(resp http.ResponseWriter, req *http.Request) {
	userID, _ := userIDFromCtx(req.Context())
	vars := mux.Vars(req)
	taskID := vars["id"]

	err := s.Store.DeleteTask(
		req.Context(),
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: taskID,
			Valid:  true,
		},
	)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
}

func (s *ToDoService) deleteTasks(resp http.ResponseWriter, req *http.Request) {
	userID, _ := userIDFromCtx(req.Context())
	createdDate := value(req, "created_date")
	err := s.Store.DeleteTasks(
		req.Context(),
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		createdDate,
	)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
}

// Validate function, which will be called for each request
func (s *ToDoService) Validate(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(resp, req)
	})
}

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func (s *ToDoService) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	return req, true
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
