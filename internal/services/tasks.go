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
	"github.com/manabie-com/togo/internal/storages"
)

type Config struct {
	Port int
	Env  string
	Db   struct {
		Dsn string
	}
	Jwt struct {
		Secret string
	}
}

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Config Config
	//pointer to logger standard library
	Logger *log.Logger
	Models storages.Models
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	email := value(req, "email")
	if !s.Models.DB.ValidateUser(email, value(req, "password")) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect email/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(email.String)
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
	var ok bool
	req, ok = s.validToken(req)
	if !ok {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}
	var user storages.User
	email, _ := userIDFromCtx(req.Context())
	m, err := s.Models.DB.GetUserFromEmail(email)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	} else {
		user = *m
		tasks, err := s.Models.DB.RetrieveTasks(
			sql.NullString{
				//convert user id to string
				String: strconv.Itoa(user.ID),
				Valid:  true,
			},
			value(req, "created_at"),
		)
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
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	var ok bool
	req, ok = s.validToken(req)
	if !ok {
		resp.WriteHeader(http.StatusUnauthorized)
		return
	}
	var t storages.Task
	resp.Header().Set("Content-Type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&t)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	now := time.Now()
	var user storages.User
	email, _ := userIDFromCtx(req.Context())
	m, err := s.Models.DB.GetUserFromEmail(email)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	} else {
		user = *m
		t.UserID = user.ID
		t.CreatedAt = now
		t.UpdatedAt = now
		t.CreatedAt = now
		t.UpdatedAt = now
		lastInsertId, err := s.Models.DB.AddTask(t, user)
		if err != nil {
			resp.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}
		t.ID = lastInsertId
		json.NewEncoder(resp).Encode(map[string]*storages.Task{
			"data": &t,
		})
	}

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
