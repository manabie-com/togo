package storages

// User reflects users data from DB
type User struct {
	ID       string `json:"user_id" validate:"required"`
	Password string `json:"password" validate:"required"`
	MaxTodo  int    `json:"-"`
}
