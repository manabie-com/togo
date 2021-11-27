package service

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"main/config"
	"main/internal/logger"
	"main/internal/model"
	"main/internal/store"
	"net/http"
)

type Service interface {
	CreateTodoHandler() http.Handler
}

type TogoService struct {
	cfg   config.Config
	store store.Store
	log   logger.Logger
}

func NewTogoService(cfg config.Config, store store.Store, log logger.Logger) (*TogoService, error) {
	return &TogoService{
		cfg:   cfg,
		store: store,
		log:   log,
	}, nil
}

type CreateTodoRequest struct {
	Title  string `json:"title"`
	UserId uint   `json:"user_id"`
}

type CreateTodoResponse struct {
	Todo model.Todo `json:"todo"`
}

func (s *TogoService) CreateTodoHandler() http.Handler {
	fn := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			http.Error(w, "Method unsupported", 405)
			return
		}

		objectReq, err := s.createTodoReqParser(req)
		if err != nil {
			http.Error(w, "Bad Request", 400)
			return
		}

		objectResp, err := s.CreateTodo(objectReq)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		resp, err := s.createTodoRespParser(objectResp)
		if err != nil {
			http.Error(w, "Internal Server Error", 500)
			return
		}

		_, err = w.Write(resp)
		if err != nil {
			s.log.Error("can not response", logger.Error(err))
			return
		}
	}
	return http.HandlerFunc(fn)
}

func (s *TogoService) createTodoReqParser(req *http.Request) (*CreateTodoRequest, error) {
	var createTodoReq *CreateTodoRequest
	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &createTodoReq)
	if err != nil {
		return nil, err
	}

	if createTodoReq == nil {
		return nil, errors.New("empty request")
	}

	if createTodoReq.UserId == 0 || createTodoReq.Title == "" {
		return nil, errors.New("missing required fields")
	}

	return createTodoReq, err
}

func (s *TogoService) CreateTodo(req *CreateTodoRequest) (*CreateTodoResponse, error) {
	todo, err := s.store.CreateTodo(model.Todo{
		Title:  req.Title,
		UserId: req.UserId,
	})

	if err != nil {
		return nil, err
	}

	return &CreateTodoResponse{Todo: todo}, nil
}

func (s *TogoService) createTodoRespParser(req *CreateTodoResponse) ([]byte, error) {
	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	return body, err
}
