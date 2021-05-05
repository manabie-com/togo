package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  *sqllite.LiteDB
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")
	resp.Header().Set("Content-Type", "application/json")
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
	var (
		err error
		token string
		statusCode int
	)
	user := &storages.User{}
	err = json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		log.Printf("error while decoding body - %s", err.Error())
		goto ERROR
	}
	if !s.Store.ValidateUser(req.Context(),
		NullString(user.ID), NullString(user.Password)) {
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
		req.Context(),
		NullString(id),
		NullString(req.FormValue("created_date")),
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
			log.Printf("error while closing body - %s", err.Error())
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

	isReachLimit, err := s.Store.IsTaskReachLimit(req.Context(), NullString(t.UserID), NullString(t.CreatedDate))
	if err != nil {
		errorResp(resp, err, 0)
		return
	}

	if isReachLimit {
		response(resp, http.StatusUnauthorized, map[string]string{
			"message": "max limit tasks reached",
		})
		return
	}

	err = s.Store.AddTask(req.Context(), t)
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
		log.Printf("error while encoding response's message - %s", err.Error())
	}
}

func NullString(value string) sql.NullString {
	return sql.NullString{
		String: value,
		Valid:  true,
	}
}
