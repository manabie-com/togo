package handlers

import "time"

type Response struct {
	Code int8        `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type CreateTaskRequest struct {
	Email string `json:"email"`
	Task  string `json:"task"`
}

type GetTaskQuery struct {
	Email string `json:"email"`
}

type SetConfigRequest struct {
	Email string    `json:"email"`
	Limit int8      `json:"limit"`
	Date  string 	`json:"date"`
	date  time.Time
}
