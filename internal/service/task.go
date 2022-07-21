package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tonghia/togo/internal/store"
)

func (s *Service) insertTask(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	userIDParam := params["userID"]
	userID, _ := strconv.ParseUint(userIDParam, 10, 64)
	if userID == 0 {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

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

	// validate daily limit
	rs, err := s.store.GetTotalTaskByUserID(context.Background(), userID)
	if err != nil {
		http.Error(w, "Error getTask", http.StatusInternalServerError)
		return
	}
	userLimit := s.userLimitSvc.GetUserLimit(userID)
	if uint32(rs.TotalTask) >= userLimit {
		http.Error(w, "Forbidden: maximum daily task reached", http.StatusForbidden)
		return
	}

	// record data
	_, err = s.store.InsertTask(context.Background(), store.InsertTaskParams{
		UserID:   userID,
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
	params := mux.Vars(req)
	userIDParam := params["userID"]
	userID, _ := strconv.ParseUint(userIDParam, 10, 64)
	if userID == 0 {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	tasks, err := s.store.GetTaskByUserID(context.Background(), userID)
	if err != nil {
		http.Error(w, "Error getTask", http.StatusInternalServerError)
		return
	}

	t := GetTodoTaskResponse{Message: "OK", Data: tasks}
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
