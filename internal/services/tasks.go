package services

import (
	"context"
	"database/sql"
	"encoding/json"
	usersqlstore"github.com/manabie-com/togo/internal/storages/user/sqlstore"
	tasksqlstore"github.com/manabie-com/togo/internal/storages/task/sqlstore"
	"github.com/manabie-com/togo/pkg/common/crypto"
	"github.com/manabie-com/togo/up"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	usermodel "github.com/manabie-com/togo/internal/storages/user/model"
	taskmodel "github.com/manabie-com/togo/internal/storages/task/model"
)

// ToDoService implement HTTP server
type ToDoService struct {
	jwtKey string
	userstore *usersqlstore.UserStore
	taskstore *tasksqlstore.TaskStore
}

func NewToDoService(db *sql.DB, JWTKey string) *ToDoService {
	return &ToDoService{
		jwtKey:    JWTKey,
		userstore: usersqlstore.NewUserStore(db),
		taskstore: tasksqlstore.NewTaskStore(db),
	}
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
	case "/register":
		if req.Method != http.MethodPost {
			return
		}
		s.Register(resp, req)
	case "/login":
		if req.Method != http.MethodPost {
			return
		}
		s.Login(resp, req)
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
			s.ListTasks(resp, req)
		case http.MethodPost:
			s.AddTask(resp, req)
		}
		return
	}
}

func (s *ToDoService) Register(resp http.ResponseWriter, req *http.Request) {
	u := &up.RegisterRequest{}
	err := json.NewDecoder(req.Body).Decode(u)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if u.MaxTodo == 0 {
		u.MaxTodo = 5
	}


	userID := sql.NullString{
		String: u.ID,
		Valid:  true,
	}
	user, err := s.userstore.FindByID(req.Context(), userID)
	if err != nil {
		if err != sql.ErrNoRows {
			resp.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if user != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "user exists with the given id",
		})
		return
	}

	err = s.userstore.Create(req.Context(), &usermodel.User{
		ID:       u.ID,
		Password: u.Password,
		MaxTodo:  u.MaxTodo,
	})
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}


	resp.Header().Set("Content-Type", "application/json")
	registerResponse := &up.RegisterResponse{
		ID:      u.ID,
		MaxTodo: u.MaxTodo,
	}

	json.NewEncoder(resp).Encode(map[string]interface{}{
		"data": registerResponse,
	})
}

func (s *ToDoService) Login(resp http.ResponseWriter, req *http.Request) {
	u := &usermodel.User{}
	err := json.NewDecoder(req.Body).Decode(u)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID := sql.NullString{
		String: u.ID,
		Valid:  true,
	}

	user, err := s.userstore.FindByID(req.Context(), userID)
	if err != nil {
		if err == sql.ErrNoRows {
			resp.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": "incorrect user_id/pwd",
			})
			return
		}
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}

	if user == nil || !crypto.CheckPasswordHash(u.Password, user.Password) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(u.ID)
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

func (s *ToDoService) ListTasks(resp http.ResponseWriter, req *http.Request) {
	id, _ := userIDFromCtx(req.Context())
	tasks, err := s.taskstore.RetrieveTasks(
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

	json.NewEncoder(resp).Encode(map[string][]*taskmodel.Task{
		"data": tasks,
	})
}

func (s *ToDoService) AddTask(resp http.ResponseWriter, req *http.Request) {
	t := &taskmodel.Task{}
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

	err = s.taskstore.AddTask(req.Context(), t)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*taskmodel.Task{
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
	token, err := at.SignedString([]byte(s.jwtKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (s *ToDoService) validToken(req *http.Request) (*http.Request, bool) {
	token := req.Header.Get("Authorization")

	claims := make(jwt.MapClaims)
	t, err := jwt.ParseWithClaims(token, claims, func(*jwt.Token) (interface{}, error) {
		return []byte(s.jwtKey), nil
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
