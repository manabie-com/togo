package tests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type registerReq struct {
	FullName    string `json:"fullName"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	TasksPerDay int    `json:"tasksPerDay"`
}

func sendRegister(ctx context.Context, req *registerReq) error {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return err
	}
	r, err := http.NewRequest("POST", host+"/auth/register", bytes.NewBuffer(reqBody))
	if err != nil {
		return err
	}
	r.Header.Set("Content-Type", "application/json")
	resp, err := new(http.Client).Do(r)
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("Wrong status: %v", resp.StatusCode)
	}
	if err != nil {
		return err
	}
	return nil
}

type loginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type loginRes struct {
	Token string `json:"token"`
}

func sendLogin(ctx context.Context, req *loginReq) (*loginRes, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", host+"/auth/login", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	resp, err := new(http.Client).Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status: %v", resp.StatusCode)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	res := new(loginRes)
	if err := json.Unmarshal(resBody, res); err != nil {
		return nil, err
	}
	return res, nil
}

type addTaskReq struct {
	Content string `json:"content"`
}

type taskData struct {
	ID      uint   `json:"id"`
	UserID  uint   `json:"userId"`
	Content string `json:"content"`
}

func sendAddTask(ctx context.Context, authToken string, req *addTaskReq) (*taskData, error) {
	reqBody, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}
	r, err := http.NewRequest("POST", host+"/tasks", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}
	r.Header.Set("Content-Type", "application/json")
	if authToken != "" {
		r.Header.Set("Authorization", "Bearer "+authToken)
	}
	resp, err := new(http.Client).Do(r)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusCreated {
		return nil, fmt.Errorf("Wrong status: %v", resp.StatusCode)
	}
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	res := new(taskData)
	if err := json.Unmarshal(resBody, res); err != nil {
		return nil, err
	}
	return res, nil
}

func sendGetTasks(ctx context.Context) ([]*taskData, error) {
	resp, err := http.Get(host + "/tasks")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Wrong status: %v", resp.StatusCode)
	}
	res := make([]*taskData, 0)
	if err := json.Unmarshal(resBody, &res); err != nil {
		return nil, err
	}
	return res, nil
}
