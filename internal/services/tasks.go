package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/manabie-com/togo/internal/storages"
	"github.com/manabie-com/togo/internal/storages/redis"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	Store  *gorm.DB
}

//var ctx = context.Background()

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
		switch req.Method {
		case http.MethodPost:
			s.getAuthToken(resp, req)
		}
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
			// s.listTasks(resp, req)
		case http.MethodPost:
			s.addTask(resp, req)
		}
		return
	case "/ping":
		s.ping(resp, req)
		return
	}
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	auth := storages.GetLogin()
	user := storages.GetUser()
	err := json.NewDecoder(req.Body).Decode(&auth)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.Store.Where("id = ? and password = ?",  auth.UserID, auth.Password).Find(&user).Error
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}

	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(user.ID)
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

//func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request) {
//	id, _ := userIDFromCtx(req.Context())
//	tasks, err := s.Store.RetrieveTasks(
//		req.Context(),
//		sql.NullString{
//			String: id,
//			Valid:  true,
//		},
//		value(req, "created_date"),
//	)
//
//	resp.Header().Set("Content-Type", "application/json")
//
//	if err != nil {
//		resp.WriteHeader(http.StatusInternalServerError)
//		json.NewEncoder(resp).Encode(map[string]string{
//			"error": err.Error(),
//		})
//		return
//	}
//
//	json.NewEncoder(resp).Encode(map[string][]*storages.Task{
//		"data": tasks,
//	})
//}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request) {
	ct := storages.GetCreateTask()
	err := json.NewDecoder(req.Body).Decode(&ct)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	t := storages.GetTask()
	now := time.Now()
	userID, _ := userIDFromCtx(req.Context())
	t.ID = uuid.New().String()
	t.UserID = userID
	t.CreatedAt = now
	t.UpdatedAt = now
	t.Content = ct.Content

	err = validationLimitedToCreate(req.Context(), &t, 5)
	if err != nil {
		resp.WriteHeader(http.StatusForbidden)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	err = s.Store.Create(&t).Error
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": &t,
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

func validationLimitedToCreate(ctx context.Context, t *storages.Task, n int64) error {
	client := redis.Init()
	year, month, day := time.Now().Date()
	sYear := strconv.Itoa(year)
	sMonth := strconv.Itoa(int(month))
	sDay := strconv.Itoa(day)
	key := t.UserID+sYear+sMonth+sDay

	err := client.SetNX(ctx, key, 0)
	if err != nil {
		return err
	}

	integer, err := client.Incr(ctx, key)
	if err != nil {
		return err
	}
	if integer > n {
		return errors.New("Create only 5 tasks only per day")
	}

	return nil
}
