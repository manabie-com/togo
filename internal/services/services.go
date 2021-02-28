package services

import (
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/services/helper"
	tasksqlstore "github.com/manabie-com/togo/internal/storages/task/sqlstore"
	usersqlstore "github.com/manabie-com/togo/internal/storages/user/sqlstore"
	"log"
	"net/http"
	"time"
)

// ToDoService implement HTTP server
type ToDoService struct {
	jwtKey    string
	userstore *usersqlstore.UserStore
	taskstore *tasksqlstore.TaskStore
	maxTodo   int

	mapUserAndTodos map[string]*helper.ToDos
}

func NewToDoService(
	db *sql.DB, JWTKey string,
	maxTodo int) *ToDoService {
	return &ToDoService{
		jwtKey:          JWTKey,
		userstore:       usersqlstore.NewUserStore(db),
		taskstore:       tasksqlstore.NewTaskStore(db),
		maxTodo:         maxTodo,
		mapUserAndTodos: make(map[string]*helper.ToDos),
	}
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
	case "/register":
		if req.Method != http.MethodPost {
			return
		}
		s.Register(resp, req)
	case "/login":
		if req.Method != http.MethodPost {
			return
		}
		s.Login(resp, req)
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
			s.ListTasks(resp, req)
		case http.MethodPost:
			s.AddTask(resp, req)
		}
		return
	}
}

// support for test
func (s *ToDoService) ClearMapUserAndTodos() {
	s.mapUserAndTodos = make(map[string]*helper.ToDos)
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
	token, err := at.SignedString([]byte(s.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.jwtKey), nil
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
