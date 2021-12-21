package model

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedDate string `json"createdDate"`
	UserID      string `json:"userId"`
}

type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
