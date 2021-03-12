package entities

// User reflect User entity in general which is decoupled from
// database entity
type User struct {
	ID       string `json:"user_id"`
	Password string `json:"password"`
}
