package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/model"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/storages"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  interface {
		RetrieveTasks(ctx context.Context, userID, createdDate sql.NullString) ([]*storages.Task, error)
		AddTask(ctx context.Context, t *storages.Task) error
		ValidateUser(ctx context.Context, userID, pwd sql.NullString) bool
		GetUserInfo(ctx context.Context, userID sql.NullString) *storages.User
		CountTasks(ctx context.Context, userID, createdDate sql.NullString) (int, error)
	}
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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
	id := value(req, "user_id")
	if !s.Store.ValidateUser(req.Context(), id, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(model.LoginResponse{
			Error: "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id.String)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(model.LoginResponse{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(model.LoginResponse{
		Data: token,
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
		json.NewEncoder(resp).Encode(model.GetTaskResponse{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(model.GetTaskResponse{
		Data: tasks,
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
	userID, _ := userIDFromCtx(req.Context())
	t.UserID = userID
	t.CreatedDate = now.UTC().Format("2006-01-02")
	t.ID = uuid.New().String()

	pUserID := sql.NullString{
		String: userID,
		Valid:  true,
	}
	pCreatedDate := sql.NullString{
		String: t.CreatedDate,
		Valid:  true,
	}

	// Retrieve user's max_todo
	userInfo := s.Store.GetUserInfo(req.Context(), pUserID)
	if userInfo == nil {
		// The token has been passed the validation
		// UserID should be existed in data store
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Count task created today
	total, err := s.Store.CountTasks(req.Context(), pUserID, pCreatedDate)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Check for the limited per day
	if total >= userInfo.MaxTodo {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(model.CreateTaskResponse{
			Error: fmt.Sprintf("You have created %d tasks today", total),
		})
		return
	}

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(model.CreateTaskResponse{
			Error: err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(model.CreateTaskResponse{
		Data: t,
	})
}

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func (s *ToDoService) CreateTokenWithExpireTime(id string, t time.Duration) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(t).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoService) createToken(id string) (string, error) {
	return s.CreateTokenWithExpireTime(id, time.Minute*15)
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
