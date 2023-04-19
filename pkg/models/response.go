package models

type Response[T any] struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Data    []T    `json:"data"`
}

type TaskRequest struct {
	Content string `json:"content"`
}

type AuthRequest struct {
	UserName string `json:"userName"`
	PassWord string `json:"PassWord"`
}

type LoginResponse struct {
	Message string `json:"message"`
	Status  int64  `json:"status"`
	Token   string `json:"token"`
}
