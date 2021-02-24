package model

// User reflects users data from DB
type User struct {
	ID       string `json:"user_id"`
	Password string `json:"password"`
}
