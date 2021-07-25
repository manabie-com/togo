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
	_ "github.com/lib/pq"
	"github.com/surw/togo/internal/storages"
)

func NewToDoService(db ILiteDB) *ToDoService {
	return &ToDoService{
		JWTKey: "wqGyEBBfPK9w3Lxw",
		Store: db,
		limitController: newLimitController(),
	}
}

func (s *ToDoService) Register(router *Router) {
	defaultInterceptors := NewInterceptor(s.logInterceptor())
	withAuthInterceptors := NewInterceptor(s.logInterceptor(), s.authInterceptor())
	router.AddHandler("/login", s.getAuthToken, defaultInterceptors)
	router.AddHandler("/tasks", s.listTasks, withAuthInterceptors, "GET")
	router.AddHandler("/tasks", s.addTask, withAuthInterceptors, "POST")
}

type ILiteDB interface {
	RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error)
	AddTask(ctx context.Context, t *storages.Task) error
	ValidateUser(ctx context.Context, userID, pwd string) error
}

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey          string
	Store           ILiteDB
	limitController limitController
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
	id := req.FormValue("user_id")
	if err = s.Store.ValidateUser(req.Context(), id, req.FormValue("password")); err != nil {
		return nil, newError(http.StatusUnauthorized, err.Error())
	}

	token, err := s.createToken(id)
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
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = todayStr

	fallbackFn, err := s.limitController.ReachLimit(req.Context(), t.UserID, t.CreatedDate, 5)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			fallbackFn()
		}
	}()
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
