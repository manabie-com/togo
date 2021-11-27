package model

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	UserId uint   `json:"user_id"`
}
