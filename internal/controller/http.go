package controller

import (
	"context"
	"encoding/json"
	"github.com/manabie-com/togo/internal/entities"
	"github.com/manabie-com/togo/internal/transport"
	"github.com/manabie-com/togo/internal/usecase"
	"log"
	"net/http"
	"time"
)

// ToDoService implement HTTP server
type ToDoService struct {
	JWTKey  string
	UseCase usecase.UseCase      `inject:""`
	Trans   *transport.Transport `inject:""`
}

func NewToDoService(db string) *ToDoService {
	todo := &ToDoService{
		JWTKey:  "wqGyEBBfPK9w3Lxw",
		Trans:   transport.NewTransport(),
		UseCase: usecase.NewUc(db),
	}
	return todo
}

func (s *ToDoService) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Println(req.Method, req.URL.Path)
	switch req.URL.Path {
	case "/login":
		s.getAuthToken(resp, req)
		return
	case "/tasks":
		token := s.Trans.GetToken(req)
		id, ok := s.UseCase.ValidToken(token, s.JWTKey)
		req = req.WithContext(context.WithValue(req.Context(), transport.UserAuthKey(0), id))
		if !ok {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}

		/*if !s.UseCase.ValidToken(token, s.JWTKey) {
			resp.WriteHeader(http.StatusUnauthorized)
			return
		}*/

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
	user := s.Trans.GetValue(req, "user_id")
	pass := s.Trans.GetValue(req, "password")
	if !s.UseCase.Validate(req.Context(), user, pass) {
		resp.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": "incorrect user_id/pwd",
		})
		return
	}
	resp.Header().Set("Content-Type", "application/json")

	token, err := s.UseCase.CreateToken(user.String, s.JWTKey)
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
	id := s.Trans.GetUserIDFromCtx(req.Context())
	date := s.Trans.GetValue(req, "created_date")
	tasks, err := s.UseCase.List(req.Context(), id, date.String)

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
	task, err := s.Trans.GetTask(req)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	id := s.Trans.GetUserIDFromCtx(req.Context())
	err = s.UseCase.Add(req.Context(), id, time.Now().Format("2006-01-02"), task)
	if err != nil {
		if err.Error() == "quantity exceeded allowed" {
			resp.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(resp).Encode(map[string]string{
				"error": err.Error(),
			})
			return
		}
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(map[string]string{
			"error": err.Error(),
		})
		return
	}

	json.NewEncoder(resp).Encode(map[string]*entities.Task{
		"data": task,
	})
}
