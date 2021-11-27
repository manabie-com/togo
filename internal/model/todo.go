package model

type Todo struct {
	Id     uint   `json:"id"`
	Title  string `json:"title"`
	UserId uint   `json:"user_id"`
}
