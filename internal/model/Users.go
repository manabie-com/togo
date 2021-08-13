package model

type Users struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	MaxTodo string  `json:"max_todo"`
}

type UserList []Users