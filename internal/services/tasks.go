package services

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/entities"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  storages.Store
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
	id := valueString(req, "user_id")
	if !s.Store.ValidateUser(req.Context(), id, valueString(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id.String)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]string{
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
		valueString(req, "created_date"),
		valueInt(req, "total"),
		valueInt(req, "page"),
	)

	resp.Header().Set("Content-Type", "application/json")

	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string][]*entities.Task{
		"data": tasks,
	})
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &entities.Task{}
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

	json.NewEncoder(resp).Encode(map[string]*entities.Task{
		"data": t,
	})
}

func valueString(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func valueInt(req *http.Request, p string) sql.NullInt32 {
	var num int32
	i, err := strconv.Atoi(req.FormValue(p))
	if err != nil {
		num = 1
		if p == "total" {
			num = 5
		}
	} else {
		num = int32(i)
	}

	return sql.NullInt32{
		Int32: num,
		Valid: true,
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
