package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/google/martian/log"
	"github.com/manabie-com/togo/internal/models"
	"io"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
)

var (
	MaxLimitReach = errors.New("max limit tasks reached")
	InvalidUserPwd = errors.New("invalid username/pwd")
)

type ResponseData struct {
	Data interface{} `json:"data"`
}

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  storages.IDatabase
	cache  ICache
	maxTodo int32
}

func NewTodoService(db storages.IDatabase, jwtToken string, pool *redis.Pool, maxTodo int32) *ToDoService {
	return &ToDoService{
		JWTKey: jwtToken,
		Store:  db,
		cache: &RedisCache{redisPool: pool},
		maxTodo: maxTodo,
	}
}

func (s *ToDoService) AddHandler(api *API) {
	api.Router.HandleFunc(LOGIN, s.GetAuthToken).Methods(http.MethodPost)
	api.Router.HandleFunc(TASKS, s.AddTask).Methods(http.MethodPost)
	api.Router.HandleFunc(TASKS, s.ListTasks).Methods(http.MethodGet)
	api.Router.HandleFunc(SIGNUP, s.SignUp).Methods(http.MethodPost)
}

func (s *ToDoService) getUser(body io.Reader) (*models.User, error) {
	user := &models.User{}
	err := json.NewDecoder(body).Decode(user)
	return user, err
}

func (s *ToDoService) SignUp(resp http.ResponseWriter, req *http.Request) {
	var token string
	user, err := s.getUser(req.Body)
	if err != nil {
		log.Errorf("error while getting body from request - %s", err.Error())
		goto ERROR
	}
	// TODO: do validate userId and password
	// store user to database
	if err = s.Store.AddUser(user.ID, user.Password, s.maxTodo); err != nil {
		goto ERROR
	}
	token, err = s.createToken(user.ID)
	if err != nil {
		goto ERROR
	}
	response(resp, 0, token)
	return
	ERROR:
	errorResp(resp, err, 0)
	return
}

func (s *ToDoService) GetAuthToken(resp http.ResponseWriter, req *http.Request) {
	var (
		err error
		token string
		statusCode int
		user *models.User
	)
	user, err = s.getUser(req.Body)
	if err != nil {
		log.Errorf("error while getting body from request - %s", err.Error())
		goto ERROR
	}
	if !s.Store.ValidateUser(user.ID, user.Password) {
		err = InvalidUserPwd
		statusCode = http.StatusBadRequest
		goto ERROR
	}
	token, err = s.createToken(user.ID)
	if err != nil {
		goto ERROR
	}
	response(resp, 0, token)
	return
ERROR:
	errorResp(resp, err, statusCode)
}

func (s *ToDoService) ListTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.Store.RetrieveTasks(id, req.FormValue("created_date"))
	if err != nil {
		errorResp(resp, err, 0)
		return
	}
	response(resp, 0, tasks)
}

func (s *ToDoService) AddTask(resp http.ResponseWriter, req *http.Request) {
	t := &models.Task{}
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
	maxTodo, err := s.getMaxTodo(t.UserID)
	if err != nil {
		log.Errorf("error while getting maxtodo from userId:%s - %s", t.UserID, err.Error())
		errorResp(resp, err, 0)
		return
	}

	// get numberOfTask from redis
	numberOfTask, err := s.getNumberOfTasks(t.UserID, t.CreatedDate)
	if err != nil {
		errorResp(resp, err, 0)
		return
	}

	if maxTodo <= numberOfTask {
		response(resp, http.StatusBadRequest, MaxLimitReach.Error())
		return
	}

	// add task with callback increase number of task in cache
	err = s.Store.AddTask(t, s.cache.IncTask)
	if err != nil {
		errorResp(resp, err, 0)
		return
	}
	response(resp, 0, t)
}

func (s *ToDoService) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["id"] = id
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoService) getNumberOfTasks(userId, createdDate string) (int32, error) {
	result, err := s.cache.GetNumberOfTasks(userId, createdDate)
	if err != nil {
		goto ERROR
	}
	if result == -1 {
		// get numberOfTasks from database
		result, err = s.Store.CountTasks(userId, createdDate)
		if err != nil {
			goto ERROR
		}
		// store to cache
		err = s.cache.SetNumberOfTasks(userId, createdDate, result)
	}
	return result, nil
	ERROR:
	return -1, err
}

func (s *ToDoService) getMaxTodo(userId string) (int32, error) {
	result, err := s.cache.GetMaxTodo(userId)
	if err != nil {
		return -1, err
	}
	if result == -1 {
		// getMaxTodo from database
		result, err = s.Store.GetMaxTodo(userId)
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
	response(resp, statusCode, err.Error())
	return
}

func response(resp http.ResponseWriter, code int, message interface{}) {
	if code == 0 {
		code = http.StatusOK
	}
	resp.WriteHeader(code)
	err := json.NewEncoder(resp).Encode(&ResponseData{Data: message})
	if err != nil {
		log.Errorf("error while encoding response's message - %s", err.Error())
	}
}

