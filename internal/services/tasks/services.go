package tasks

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/manabie-com/togo/internal/storages"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey string
	repo   storages.Repository
}

func SetupNewService(jwtKey string, r storages.Repository) *ToDoService {
	return &ToDoService{JWTKey: jwtKey, repo: r}
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
		req, ok = s.ValidToken(req)
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		switch req.Method {
		case http.MethodDelete:
			s.deleteTaskByDate(resp, req)
		}
		return
	}
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	username := value(req, "user_name")
	user, err := s.repo.ValidateUser(username, value(req, "password"))
	if err != nil {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}

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

func (s *ToDoService) ListTasks(context context.Context, id string, created_date sql.NullString) ([]*storages.Task, error) {
	tasks, err := s.repo.RetrieveTasks(
		context,
		sql.NullString{
			String: id,
			Valid:  true,
		},
		created_date,
	)

	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (s *ToDoService) AddTask(ctx context.Context, id sql.NullString, t storages.Task) (*storages.Task, error) {
	// err := json.NewDecoder(req.Body).Decode(t)
	// defer req.Body.Close()
	// if err != nil {
	// 	resp.WriteHeader(http.StatusInternalServerError)
	// 	return
	// }

	createdDateInSqlNullString := convertStringToSqlNullString(time.Now().Format("2006-01-02"))

	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	tasks, err := s.repo.RetrieveTasks(
		ctx,
		id,
		createdDateInSqlNullString,
	)
	if err != nil {
		return nil, err
	}
	if len(tasks) == int(user.MaxTodo) {
		return nil, errors.New("exceed today maximum allowed number of tasks")
	}

	t.UserID = id.String
	t.CreatedDate = createdDateInSqlNullString.String

	taskId, err := s.repo.AddTask(ctx, &t)
	if err != nil {
		return nil, err
	}

	t.ID = taskId

	return &t, nil
}

func (s *ToDoService) deleteTaskByDate(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	createdDate := value(req, "created_date")
	userID, _ := s.UserIDFromCtx(req.Context())

	user, err := s.repo.GetUserById(req.Context(), convertStringToSqlNullString(userID))
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}
	if user == nil {
		resp.WriteHeader(http.StatusNotFound)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "User not found",
		})
		return
	}

	err = s.repo.DeleteTaskByDate(
		req.Context(),
		convertStringToSqlNullString(userID),
		createdDate,
	)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	resp.WriteHeader(http.StatusNoContent)
}

func value(req *http.Request, p string) sql.NullString {
	return sql.NullString{
		String: req.FormValue(p),
		Valid:  true,
	}
}

func convertStringToSqlNullString(s string) sql.NullString {
	return sql.NullString{
		String: s,
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

func (s *ToDoService) ValidToken(req *http.Request) (*http.Request, bool) {
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

func (s *ToDoService) UserIDFromCtx(ctx context.Context) (string, bool) {
	v := ctx.Value(userAuthKey(0))
	id, ok := v.(string)
	return id, ok
}
