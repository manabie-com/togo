package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"togo/internal/storages"
	pg "togo/internal/storages/postgres"
	sqllite "togo/internal/storages/sqlite"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

var limitNtaskPerday = make(map[string]int, 0)
var maxLimitNtaskPerday int = 5

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey  string
	Store   *sqllite.LiteDB
	StorePg *pg.ProstgresDB
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
	case "/login-pg":
		s.getAuthTokenPg(resp, req)
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
			ok = s.limitNtaskToday(resp, req)
			if !ok {
				return
			}
			s.addTask(resp, req)
		}
		return
	case "/tasks-pg":
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodGet:
			s.listTasksPg(resp, req)
		case http.MethodPost:
			ok = s.limitNtaskToday(resp, req)
			if !ok {
				return
			}
			s.addTaskPg(resp, req)
		}
		return
	}
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	if !s.Store.ValidateUser(req.Context(), id, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		formatOutput(resp, map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id.String)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		formatOutput(resp, map[string]string{
			"error": err.Error(),
		})
		return
	}

	formatOutput(resp, map[string]string{
		"data": token,
	})
}

func (s *ToDoService) getAuthTokenPg(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	if !s.StorePg.ValidateUser(req.Context(), id, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		formatOutput(resp, map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id.String)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		formatOutput(resp, map[string]string{
			"error": err.Error(),
		})
		return
	}

	formatOutput(resp, map[string]string{
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
		formatOutput(resp, map[string]string{
			"error": err.Error(),
		})
		return
	}

	formatOutput(resp, map[string][]*storages.Task{
		"data": tasks,
	})
}

func (s *ToDoService) listTasksPg(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.StorePg.RetrieveTasks(
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
		formatOutput(resp, map[string]string{
			"error": err.Error(),
		})
		return
	}

	formatOutput(resp, map[string][]*storages.Task{
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

	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = getCurrentDate()

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		formatOutput(resp, map[string]string{
			"error": err.Error(),
		})
		return
	}

	key := getCurrentDate() + "_" + userID
	limitNtaskPerday[key]++
	formatOutput(resp, map[string]*storages.Task{
		"data": t,
	})
}

func (s *ToDoService) addTaskPg(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = getCurrentDate()

	resp.Header().Set("Content-Type", "application/json")

	err = s.StorePg.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		formatOutput(resp, map[string]string{
			"error": err.Error(),
		})
		return
	}

	key := getCurrentDate() + "_" + userID
	limitNtaskPerday[key]++
	formatOutput(resp, map[string]*storages.Task{
		"data": t,
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
	return at.SignedString([]byte(s.JWTKey))
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil || t == nil {
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

func formatOutput(resp http.ResponseWriter, v interface{}) {
	json.NewEncoder(resp).Encode(v)
}

func getCurrentDate() string {
	now := time.Now()
	return now.Format("2006-01-02")
}

func (s *ToDoService) limitNtaskToday(resp http.ResponseWriter, req *http.Request) bool {
	id, ok := userIDFromCtx(req.Context())
	if !ok {
		resp.WriteHeader(http.StatusBadRequest)
		formatOutput(resp, map[string]string{
			"error": "unknow",
		})
		return false
	}

	key := getCurrentDate() + "_" + id
	if limitNtaskPerday[key] >= maxLimitNtaskPerday {
		resp.WriteHeader(http.StatusLocked)
		formatOutput(resp, map[string]string{
			"error": "created tasks over limit per day",
		})
		return false
	}
	return true
}
