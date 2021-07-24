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
	_ "github.com/mattn/go-sqlite3"
	"github.com/surw/togo/internal/storages"
	sqllite "github.com/surw/togo/internal/storages/sqlite"
)

func NewToDoService() *ToDoService {
	db, err := sql.Open("sqlite3", "./data.db")
	if err != nil {
		log.Fatal("error opening db", err)
	}
	return &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: &sqllite.LiteDB{
			DB: db,
		},
	}
}

func (s *ToDoService) Serve(port int32, router *Router) {
	defaultInterceptors := NewInterceptor(s.logInterceptor())
	withAuthInterceptors := NewInterceptor(s.logInterceptor(), s.authInterceptor())
	router.AddHandler("/login", s.getAuthToken, defaultInterceptors)
	router.AddHandler("/tasks", s.listTasks, withAuthInterceptors, "GET")
	router.AddHandler("/tasks", s.addTask, withAuthInterceptors, "POST")

	router.Start(port)
}

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  *sqllite.LiteDB
}

func (s *ToDoService) logInterceptor() httpInterceptor {
	return func(req *http.Request, handler httpHandler) (resp interface{}, err error) {
		resp, err = handler(req)
		if err != nil {
			log.Println(req.Method, req.URL.Path, err)
		}
		return
	}
}
func (s *ToDoService) authInterceptor() httpInterceptor {
	return func(req *http.Request, handler httpHandler) (resp interface{}, err error) {
		newReq, ok := s.validToken(req)
		if !ok {
			return nil, newError(http.StatusUnauthorized, "invalid token")
		}
		return handler(newReq)
	}
}

func (s *ToDoService) getAuthToken(req *http.Request) (resp interface{}, err error) {
	id := value(req, "user_id")
	if err = s.Store.ValidateUser(req.Context(), id, value(req, "password")); err != nil {
		return nil, newError(http.StatusUnauthorized, err.Error())
	}

	token, err := s.createToken(id.String)
	if err != nil {
		return nil, err
	}
	return token, nil
}
func (s *ToDoService) listTasks(req *http.Request) (resp interface{}, err error) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.Store.RetrieveTasks(
		req.Context(),
		sql.NullString{
			String: id,
			Valid:  true,
		},
		value(req, "created_date"),
	)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}
func (s *ToDoService) addTask(req *http.Request) (resp interface{}, err error) {
	t := &storages.Task{}
	err = json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		return nil, err
	}

	return t, nil
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
