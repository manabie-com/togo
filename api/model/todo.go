package model

// Todo main model of todo
type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedDate string `json"createdDate"`
	UserID      string `json:"userId"`
}

// TodoRequest request object in creating new todo
type TodoRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}
