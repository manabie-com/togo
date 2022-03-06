package dto

type TodoDto struct {
	ID      int    `json:"id"`
	Content string `json:"content"`
	UserId  uint   `json:"user_id"`
}
