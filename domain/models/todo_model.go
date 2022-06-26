package models

import "time"

type Todo struct {
	ID        int       `json:id`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Done      string    `json:"done"`
	FkUser    int       `json:"fkUser"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type NewTodo struct {
	Title  string `json:"title"`
	Text   string `json:"text"`
	Done   string `json:"done"`
	FkUser int    `json:"fkUser"`
}

type OverviewTodo struct {
	ID        int       `json:id`
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	Done      string    `json:"done"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
