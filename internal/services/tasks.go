package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
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

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	switch req.URL.Path {
	case "/login":
		s.createTokenHandler()(resp, req)
		return
	case "/tasks":
		switch req.Method {
		case http.MethodGet:
			s.authHandler(s.listTasksHandler())(resp, req)
		case http.MethodPost:
			s.authHandler(s.addTaskHandler())(resp, req)
		}
		return
	}
}

func (s *ToDoService) listTasksHandler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		resp.Header().Set("Content-Type", "application/json")

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
			resp.WriteHeader(http.StatusInternalServerError)
			respData := map[string]string{
				"error": err.Error(),
			}
			if err = json.NewEncoder(resp).Encode(respData); err != nil {
				log.Println(err)
			}
			return
		}

		respData := map[string][]*storages.Task{
			"data": tasks,
		}
		if err = json.NewEncoder(resp).Encode(respData); err != nil {
			log.Println(err)
		}
	}
}

/*func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
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
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*storages.Task{
		"data": tasks,
	})
}*/

func (s *ToDoService) addTaskHandler() http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		t := &storages.Task{}
		err := json.NewDecoder(req.Body).Decode(t)
		defer req.Body.Close()
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}

		userID, ok := userIDFromCtx(req.Context())
		if !ok {
			resp.WriteHeader(http.StatusBadRequest)
			return
		}

		t.ID = uuid.New().String()
		t.UserID = userID
		t.CreatedDate = time.Now().Format("2006-01-02")

		resp.Header().Set("Content-Type", "application/json")

		err = s.Store.AddTask(req.Context(), t)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			respData := map[string]string{
				"error": err.Error(),
			}
			if err := json.NewEncoder(resp).Encode(respData); err != nil {
				log.Println(err)
			}
			return
		}

		respData := map[string]*storages.Task{
			"data": t,
		}
		if err := json.NewEncoder(resp).Encode(respData); err != nil {
			log.Println(err)
		}
	}
}

/*func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedDate = now.Format("2006-01-02")

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": t,
	})
}*/

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func (s *ToDoService) createToken(id string) (string, error) {
	claims := jwt.MapClaims{
		authUserIdKey: id,
		authExpKey:     time.Now().Add(time.Minute * 15).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.JWTKey))
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, error) {
	authToken := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	parsedToken, err := jwt.ParseWithClaims(authToken, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.JWTKey), nil
	})
	if err != nil {
		return req, err
	}

	if !parsedToken.Valid {
		return req, authTokenIsNotValid
	}

	id, ok := claims[authUserIdKey].(string)
	if !ok {
		return req, authUserIdIsRequired
	}

	req = req.WithContext(context.WithValue(req.Context(), authUserIdKey, id))
	return req, nil
}

const (
	authUserIdKey string = "user_id"
	authExpKey           = "exp"
)

var (
	authTokenIsNotValid  = errors.New("auth token is not valid")
	authUserIdIsRequired = errors.New(fmt.Sprintf(`'%v' is required`, authUserIdKey))
)

func userIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(authUserIdKey)
	id, ok := v.(string)
	return id, ok
}
