package server

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/banhquocdanh/togo/internal/services"
	"github.com/banhquocdanh/togo/internal/storages"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"time"
)

type ToDoHttpServer struct {
	srv    *services.ToDoService
	JWTKey string
}

func NewToDoHttpServer(jwtKey string, srv *services.ToDoService) *ToDoHttpServer {
	return &ToDoHttpServer{
		srv:    srv,
		JWTKey: jwtKey,
	}

}

func (s *ToDoHttpServer) Listen(port int) error {
	log.Printf("Listen service on port :%d\n", port)
	return http.ListenAndServe(fmt.Sprintf(":%d", port), s)
}

func (s *ToDoHttpServer) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
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
		s.login(resp, req)
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

func (s *ToDoHttpServer) listTasks(resp http.ResponseWriter, req *http.Request) {
	userID, _ := userIDFromCtx(req.Context())
	createDate := req.FormValue("created_date")
	if createDate == "" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "Invalid create_date",
		})
		return
	}
	tasks, err := s.srv.ListTasks(req.Context(), userID, createDate)
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

type AddTaskRequest struct {
	Content string `json:"content"`
}

func (s *ToDoHttpServer) addTask(resp http.ResponseWriter, req *http.Request) {
	addTaskReq := &AddTaskRequest{}
	err := json.NewDecoder(req.Body).Decode(addTaskReq)
	defer req.Body.Close()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	if addTaskReq.Content == "" {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "invalid task's content",
		})
		return
	}
	userID, _ := userIDFromCtx(req.Context())

	task, err := s.srv.AddTask(req.Context(), userID, addTaskReq.Content)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*storages.Task{
		"data": task,
	})
}

func (s *ToDoHttpServer) login(resp http.ResponseWriter, req *http.Request) {
	id := req.FormValue("user_id")
	pw := req.FormValue("password")

	if !s.srv.ValidateUser(req.Context(), id, pw) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.createToken(id)
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

func (s *ToDoHttpServer) createToken(id string) (string, error) {
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

func (s *ToDoHttpServer) validToken(req *http.Request) (*http.Request, bool) {
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
