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
	"github.com/manabie-com/togo/internal/storages"
	sqllite "github.com/manabie-com/togo/internal/storages/sqlite"
)

// ToDoService implements HTTP server
type ToDoService struct {
	JWTKey string
	Store  *sqllite.LiteDB
}

// ServeHTTP listens to calls and executes corresponding action
func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		sendOKResponse(resp, nil)
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
			sendCodeResponse(resp, http.StatusUnauthorized, nil)
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

// getAuthToken generates token for validated user credentials
func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	if !s.Store.ValidateUser(req.Context(), id, value(req, "password")) {
		sendCodeResponse(
			resp,
			http.StatusUnauthorized,
			map[string]string{"error": "incorrect user_id/pwd"},
		)
		return
	}

	token, err := s.createToken(id.String)
	if err != nil {
		sendCodeResponse(
			resp,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	sendOKResponse(
		resp,
		map[string]string{"data": token},
	)
}

// listTasks returns all tasks under a user on a specified date
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

	if err != nil {
		sendCodeResponse(
			resp,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	sendOKResponse(
		resp,
		map[string][]*storages.Task{"data": tasks},
	)
}

// addTask adds provided task for the current date
func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	date := now.Format("2006-01-02")

	// Return status 400 if task count limit is already met
	if !s.addTaskAllowed(req.Context(), resp, userID, date) {
		sendCodeResponse(
			resp,
			http.StatusBadRequest,
			map[string]string{"error": "Task count limit exceeded."},
		)
		return
	}

	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		sendCodeResponse(
			resp,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = date

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		sendCodeResponse(
			resp,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return
	}

	// Return data into the API response
	sendOKResponse(
		resp,
		map[string]*storages.Task{"data": t},
	)
}

// addTaskAllowed checks whether user exceeded allowed task count
func (s *ToDoService) addTaskAllowed(context context.Context, resp http.ResponseWriter, userID, date string) bool {
	// Check number of existing tasks by user today
	tCount, err := s.Store.CountTasks(
		context,
		sql.NullString{
			String: userID,
			Valid:  true,
		},
		sql.NullString{
			String: date,
			Valid:  true,
		},
	)
	if err != nil {
		sendCodeResponse(
			resp,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return false
	}

	// Query task count limit of the user
	mCount, err := s.Store.RetrieveUserTaskLimit(
		context,
		sql.NullString{
			String: userID,
			Valid:  true,
		},
	)
	if err != nil {
		sendCodeResponse(
			resp,
			http.StatusInternalServerError,
			map[string]string{"error": err.Error()},
		)
		return false
	}

	// Return true if user's current task count is less than the task count limit
	if tCount < mCount {
		return true
	}

	return false
}

// value converts a string into a SQL Null String
func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

// createToken generates a token for the user
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

// validToken checks whether a given token is valid
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

// userIDFromCtx retrieves the userID from the request
func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}

// sendCodeResponse creates a response with indicated status code and payload
func sendCodeResponse(resp http.ResponseWriter, statusCode int, content interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(statusCode)
	if content != nil {
		json.NewEncoder(resp).Encode(content)
	}
}

// sendOKResponse creates a response with status code 200 and payload
func sendOKResponse(resp http.ResponseWriter, content interface{}) {
	resp.Header().Set("Content-Type", "application/json")
	if content != nil {
		json.NewEncoder(resp).Encode(content)
	}
}
