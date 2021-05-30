package users

type Users struct {
	ID       string `json:"id" form:"id"`
	Password string `json:"password" form:"password"`
	MaxTodo  int    `json:"max_todo" form:"max_todo"`
}

func (Users) TableName() string {
	return "users"
}
