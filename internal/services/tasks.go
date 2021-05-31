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
	postgres "github.com/manabie-com/togo/internal/storages/postgres"
)

const (
	LimitTask = 5
	StatusOpen = "open"
	StatusDone = "done"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  *postgres.DataBase
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	switch req.URL.Path {
	case "/register":
		s.register(resp, req)
		return
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
	case "/tasks/update":
		var ok bool
		req, ok = s.validToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		if req.Method == http.MethodPost {
			s.updateTask(resp, req)
		}
		return
	}
}

func (s *ToDoService) register(resp http.ResponseWriter, req *http.Request) {
	userID := value(req, "user_id")
	pwd := value(req, "password")
	// Need to define user name format and validate it
	if userID.String == "" || pwd.String == "" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "User ID or password is empty",
		})
		log.Println("Error register. Empty user ID/password")
		return
	}
	err := s.Store.AddUser(req.Context(),userID ,pwd)
	if err != nil {
		if err == postgres.ErrorUserExisted {
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
		}
		log.Printf("Error register. %v\n", err)
	}
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	id := value(req, "user_id")
	if !s.Store.ValidateUser(req.Context(), id, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Incorrect User ID or password",
		})
		log.Println("Error authentication. Failed")
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id.String)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		log.Printf("Error authentication. %v\n", err)
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
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error add task. %v\n", err)
		return
	}

	td, _ := time.Parse("2006-01-02", t.TargetDate)
	t.TargetDate = td.Format("2006-01-02")

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.Status = StatusOpen
	t.CreatedDate = now.Format("2006-01-02")
	if td.Before(now) {
		t.TargetDate = t.CreatedDate
	}

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.AddTask(req.Context(), t, LimitTask)
	if err != nil {
		if err == postgres.ErrorReachLimitTask {
			resp.WriteHeader(http.StatusExpectationFailed)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error add task. %v\n", err)
		}
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": t,
	})
}

func (s *ToDoService) updateTask(resp http.ResponseWriter, req *http.Request) {
	t := &storages.Task{}
	err := json.NewDecoder(req.Body).Decode(t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		log.Printf("Error update task. %v\n", err)
		return
	}

	if t.ID == "" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Null task ID",
		})
		log.Println("Error update task. Null task ID")
		return
	}
	
	if t.Status != StatusOpen && t.Status != StatusDone {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Wrong task status",
		})
		log.Println("Error update task. Wrong task status")
		return
	}

	td, _ := time.Parse("2006-01-02", t.TargetDate)
	t.TargetDate = td.Format("2006-01-02")

	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.UserID = userID
	if t.CreatedDate == "" {
		t.CreatedDate = now.Format("2006-01-02")
	}
	if td.Before(now) {
		t.TargetDate = t.CreatedDate
	}

	resp.Header().Set("Content-Type", "application/json")

	err = s.Store.UpdateTask(req.Context(), t, LimitTask)
	if err != nil {
		if err == postgres.ErrorReachLimitTask {
			resp.WriteHeader(http.StatusExpectationFailed)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
		} else {
			resp.WriteHeader(http.StatusInternalServerError)
			log.Printf("Error update task. %v\n", err)
		}
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
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
