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

type ToDoHttpResponseWriter struct {
	http.ResponseWriter
}

// Ideally this should be a method that mutates resp
func response(resp *http.ResponseWriter, httpStatusCode int, body interface{}) {
	// We should limit these to the lowest privileges possible (e.g only allow our client's URI)
	(*resp).Header().Set("Access-Control-Allow-Origin", "*")
	(*resp).Header().Set("Access-Control-Allow-Headers", "*")
	(*resp).Header().Set("Access-Control-Allow-Methods", "*")
	(*resp).Header().Set("Content-Type", "application/json")

	(*resp).WriteHeader(httpStatusCode)
	if body != nil {
		json.NewEncoder((*resp)).Encode(body)
	}
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)

	if req.Method == http.MethodOptions {
		response(&resp, http.StatusOK, nil)
		return
	}

	switch req.URL.Path {
	case "/login":
		ctx := req.Context()
		id := value(req, "user_id")
		pwd := value(req, "password")
		token, err := s.getAuthToken(ctx, id, pwd)

		body := make(map[string]string)
		if err != nil {
			body["error"] = err.Error()
			response(&resp, http.StatusUnauthorized, body)
		} else {
			body["token"] = token
			response(&resp, http.StatusOK, body)
		}
		return
	case "/tasks":
		// This should be a middleware to avoid repetition in other endpoints
		var err error
		req, err = s.parseUserIdToContext(req)
		if err != nil {
			response(&resp, http.StatusUnauthorized, map[string]string{"error": err.Error()})
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

func (s *ToDoService) parseUserIdToContext(req *http.Request) (*http.Request, error) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		log.Println(err)
		return req, err
	}

	if !t.Valid {
		return req, errors.New("Invalid token")
	}

	id, ok := claims["user_id"].(string)
	if !ok {
		return req, errors.New("Unable to fetch id")
	}

	req = req.WithContext(context.WithValue(req.Context(), USER_AUTH_KEY, id))
	return req, nil
}

func (s *ToDoService) getAuthToken(ctx context.Context,
	userID sql.NullString,
	password sql.NullString) (token string, err error) {
	if !s.Store.ValidateUser(ctx, userID, password) {
		err = errors.New("incorrect user_id/pwd") // Maybe we can use an enum here instead of hardcoded string
		return
	}

	token, err = s.createToken(userID.String)
	return
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
	if err != nil {
		body := map[string]string{"error": err.Error()}
		response(&resp, http.StatusInternalServerError, body)
		return
	}

	body := map[string][]*storages.Task{"data": tasks}
	response(&resp, http.StatusOK, body)
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		response(&resp, http.StatusInternalServerError, nil)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		response(&resp, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	response(&resp, http.StatusOK, map[string]*storages.Task{"data": t})
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
	token, err := at.SignedString([]byte(s.JWTKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

const USER_AUTH_KEY = "USER_AUTH_KEY"

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(USER_AUTH_KEY)
	id, ok := v.(string)
	return id, ok
}
