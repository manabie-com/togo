package up

type RegisterRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	MaxTodo  int    `json:"max_todo"`
}

type RegisterResponse struct {
	ID      string `json:"id"`
	MaxTodo int    `json:"max_todo"`
}