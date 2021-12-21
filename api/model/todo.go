package model

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	CreatedDate string `json"createdDate"`
	User        User   `json:"user"`
}

type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
