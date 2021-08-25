package services

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/manabie-com/togo/internal/common"
	"github.com/manabie-com/togo/internal/dto"
	"github.com/manabie-com/togo/internal/usecase"
)

// ToDoService implement HTTP server
type ToDoService struct {
	UserUsecase usecase.UserUsecase
	TaskUsecase usecase.TaskUsecase
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
		s.doTasks(resp, req)
		return
	default:
		resp.WriteHeader(http.StatusNotFound)
	}
}

func (s *ToDoService) getAuthToken(resp http.ResponseWriter, req *http.Request) {
	reqDTO := &dto.LoginRequestDTO{
		UserID:   req.FormValue("user_id"),
		Password: req.FormValue("password"),
	}
	respDTO, err := s.UserUsecase.Login(req.Context(), reqDTO)
	if err != nil {
		common.ResponseError(err, resp)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(respDTO)
}

func (s *ToDoService) doTasks(resp http.ResponseWriter, req *http.Request) {
	userID, err := s.validToken(req)
	if err != nil {
		common.ResponseError(err, resp)
		return
	}

	switch req.Method {
	case http.MethodGet:
		s.listTasks(resp, req, userID)
	case http.MethodPost:
		s.addTask(resp, req, userID)
	}
}

func (s *ToDoService) validToken(req *http.Request) (string, error) {
	token := req.Header.Get("Authorization")
	verifyTokenDTO := &dto.VerifyTokenRequestDTO{
		Token: token,
	}
	respDTO, err := s.UserUsecase.VerifyToken(req.Context(), verifyTokenDTO)
	if err != nil {
		return "", err
	}

	return respDTO.UserID, nil
}

func (s *ToDoService) listTasks(resp http.ResponseWriter, req *http.Request, userID string) {
	reqDTO := &dto.ListTasksRequestDTO{
		UserID:      userID,
		CreatedDate: req.FormValue("created_date"),
	}
	respDTO, err := s.TaskUsecase.ListTasks(req.Context(), reqDTO)
	if err != nil {
		common.ResponseError(err, resp)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(respDTO)
}

func (s *ToDoService) addTask(resp http.ResponseWriter, req *http.Request, userID string) {
	reqDTO := &dto.AddTaskRequestDTO{}
	err := json.NewDecoder(req.Body).Decode(reqDTO)
	defer req.Body.Close()
	if err != nil {
		log.Printf("Decode add task body error: %v\n", err)
		common.ResponseError(errors.New(common.ReasonInternalError.Code()), resp)
		return
	}

	reqDTO.UserID = userID
	respDTO, err := s.TaskUsecase.AddTask(req.Context(), reqDTO)
	if err != nil {
		common.ResponseError(err, resp)
		return
	}

	resp.Header().Set("Content-Type", "application/json")
	json.NewEncoder(resp).Encode(respDTO)
}
