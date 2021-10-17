package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/postgres"
	"github.com/manabie-com/togo/internal/usecases/task"
	"github.com/manabie-com/togo/internal/usecases/user"
	"github.com/manabie-com/togo/internal/utils"
)

// ToDoService implement HTTP server
type ToDoService struct {
	taskUseCase task.TaskUseCase
	userUseCase user.UserUseCase
}

func NewToDoService(db *sql.DB) *ToDoService {
	storeRepo := postgres.NewPostgresRepository(db)
	taskUseCase := task.NewTaskUseCase(storeRepo)
	userUseCase := user.NewUserUseCase(storeRepo)
	return &ToDoService{taskUseCase: taskUseCase, userUseCase: userUseCase}
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
	username := value(req, "username")
	resp.Header().Set("Content-Type", "application/json")

	if !s.userUseCase.ValidateUser(req.Context(), username, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect username/pwd",
		})
		return
	}

	user, err := s.userUseCase.GetUserByUsername(req.Context(), username)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	token, err := s.userUseCase.GenerateToken(user.ID, user.MaxTodo)
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
	id := userIDFromCtx(req.Context())
	log.Println(value(req, "created_date"))
	tasks, err := s.taskUseCase.ListTasks(
		req.Context(),
		id,
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

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID := userIDFromCtx(req.Context())
	maxTodo := userMaxTodoFromCtx(req.Context())

	t.ID = uuid.New().String()
	t.UserID = uint(userID)
	t.CreatedDate = now.Format("2006-01-02")

	isMaximum, err := s.taskUseCase.IsMaximumTasks(
		req.Context(),
		userID,
		utils.SqlString(t.CreatedDate),
		maxTodo,
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	if isMaximum {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "the maximum number of tasks is reached.",
		})
		return
	}

	err = s.taskUseCase.AddTask(req.Context(), t)
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

func value(req *http.Request, p string) sql.NullString {
	return utils.SqlString(req.FormValue(p))
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		log.Println(err)
		return req, false
	}

	if !t.Valid {
		return req, false
	}

	id, ok := claims["user_id"].(float64)
	if !ok {
		return req, false
	}

	max_todo, ok := claims["max_todo"].(float64)
	if !ok {
		return req, false
	}

	req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
	req = req.WithContext(context.WithValue(req.Context(), userMaxTodo(0), max_todo))

	return req, true
}

type userAuthKey int8
type userMaxTodo int8

func userIDFromCtx(ctx context.Context) uint {
	v := ctx.Value(userAuthKey(0))
	return uint(v.(float64))
}

func userMaxTodoFromCtx(ctx context.Context) uint {
	v := ctx.Value(userMaxTodo(0))
	return uint(v.(float64))
}
