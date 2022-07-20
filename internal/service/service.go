package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/tonghia/togo/internal/store"
)

type Service struct {
	store store.Querier
}

func NewService(store store.Querier) *Service {
	return &Service{store}
}

func (s *Service) Health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok")
}

func (s *Service) insertTask(w http.ResponseWriter, req *http.Request) {
	bytes, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, "Error reading body", http.StatusInternalServerError)
		return
	}
	p := CreateTodoTaskRequest{}
	err = json.Unmarshal(bytes, &p)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}

	_, err = s.store.InsertTask(context.Background(), store.InsertTaskParams{
		UserID:   p.UserID,
		TaskName: p.Name,
	})
	if err != nil {
		fmt.Printf("Error InsertTask: %v \n", err)
		http.Error(w, "Error RecordTask", http.StatusInternalServerError)
		return
	}

	t := CreateTodoTaskResponse{Message: "Success"}
	jd, err := json.Marshal(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jd)
}

func (s *Service) getTask(w http.ResponseWriter, req *http.Request) {
	task, err := s.store.GetTaskByID(context.Background(), 1) // TODO: use correct id
	if err != nil {
		http.Error(w, "Error getTask", http.StatusInternalServerError)
		return
	}

	t := GetTodoTaskResponse{Name: task.TaskName}
	jd, err := json.Marshal(&t)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jd)
}

func (s *Service) RecordTask(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "POST":
		s.insertTask(w, req)
	case "GET":
		s.getTask(w, req)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
