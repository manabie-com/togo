package models

import "time"

type User struct {
	ID         int       `json:id`
	Username   string    `json:"username"`
	DailyLimit int       `json:"dailyLimit"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

type NewUser struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	DailyLimit int    `json:"dailyLimit"`
}

type OverviewUser struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
}
