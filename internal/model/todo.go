package model

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	UserId string `json:"user_id"`
}
