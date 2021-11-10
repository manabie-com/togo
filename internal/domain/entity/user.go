package entity

// User entity
type User struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashedPassword"`
	MaxTodo        int    `json:"maxTodo"`
}
