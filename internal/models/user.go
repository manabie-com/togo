package models

// User reflects users data from DB
type User struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}
