package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/google/martian/log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  storages.IDatabase
	cache  ICache
}

func NewTodoService(db storages.IDatabase, jwtToken string, pool *redis.Pool) *ToDoService {
	return &ToDoService{
		JWTKey: jwtToken,
		Store:  db,
		cache: &RedisCache{redisPool: pool},
	}
}

func (s *ToDoService) AddHandler(api *API) {
	api.Router.HandleFunc("/login", s.getAuthToken).Methods(http.MethodPost)
	api.Router.HandleFunc("/tasks", s.addTask).Methods(http.MethodPost)
	api.Router.HandleFunc("/tasks", s.listTasks).Methods(http.MethodGet)
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	var (
		err error
		token string
		statusCode int
	)
	user := &storages.User{}
	err = json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		log.Errorf("error while decoding body - %s", err.Error())
		goto ERROR
	}
	if !s.Store.ValidateUser(req.Context(),
		user.ID, user.Password) {
		err = errors.New("incorrect user_id/pwd")
		statusCode = http.StatusUnauthorized
		goto ERROR
	}
	token, err = s.createToken(user.ID)
	if err != nil {
		goto ERROR
	}
	response(resp, 0, map[string]interface{}{
		"data": token,
	})
	return
ERROR:
	errorResp(resp, err, statusCode)
}

func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.Store.RetrieveTasks(
		req.Context(), id, req.FormValue("created_date"),
	)
	if err != nil {
		errorResp(resp, err, 0)
		return
	}
	response(resp, 0, map[string][]*storages.Task{
		"data": tasks,
	})
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer func() {
		if err := req.Body.Close(); err != nil {
			log.Errorf("error while closing body - %s", err.Error())
		}
	} ()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	// get maxtodo
	maxTodo, err := s.getMaxTodo(req.Context(), t.UserID)
	if err != nil {
		log.Errorf("error while getting maxtodo from userId:%s - %s", t.UserID, err.Error())
		errorResp(resp, err, 0)
		return
	}

	// get numberOfTask from redis
	numberOfTask, err := s.getNumberOfTasks(req.Context(), t.UserID, t.CreatedDate)
	if err != nil {
		errorResp(resp, err, 0)
		return
	}

	if maxTodo <= numberOfTask {
		response(resp, http.StatusUnauthorized, map[string]string{
			"message": "max limit tasks reached",
		})
		return
	}

	// add task with callback increase number of task in cache
	err = s.Store.AddTask(req.Context(), t, s.cache.IncTask)
	if err != nil {
		errorResp(resp, err, 0)
		return
	}
	response(resp, 0, map[string]*storages.Task{
		"data": t,
	})
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

func (s *ToDoService) getNumberOfTasks(ctx context.Context, userId, createdDate string) (int32, error) {
	result, err := s.cache.GetNumberOfTasks(userId, createdDate)
	if err != nil {
		goto ERROR
	}
	if result == -1 {
		// get numberOfTasks from database
		result, err = s.Store.CountTasks(ctx, userId, createdDate)
		if err != nil {
			goto ERROR
		}
		// store to cache
		err = s.cache.SetNumberOfTasks(userId, createdDate, result)
	}
	ERROR:
	return -1, err
}

func (s *ToDoService) getMaxTodo(ctx context.Context, userId string) (int32, error) {
	result, err := s.cache.GetMaxTodo(userId)
	if err != nil {
		return -1, err
	}
	if result == -1 {
		// getMaxTodo from database
		result, err = s.Store.GetMaxTodo(ctx, userId)
		if err != nil {
			return -1, err
		}
		// store to cache
		err = s.cache.SetMaxTodo(userId, result)
	}
	return result, err
}

type userAuthKey int8

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}

func errorResp(resp http.ResponseWriter, err error, code int) {
	statusCode := http.StatusInternalServerError
	if http.StatusText(code) != "" {
		statusCode = code
	}
	response(resp, statusCode, map[string]interface{}{
		"error": err.Error(),
	})
	return
}

func response(resp http.ResponseWriter, code int, message interface{}) {
	if code == 0 {
		code = http.StatusOK
	}
	resp.WriteHeader(code)
	err := json.NewEncoder(resp).Encode(message)
	if err != nil {
		log.Errorf("error while encoding response's message - %s", err.Error())
	}
}

