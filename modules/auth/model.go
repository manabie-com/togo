package auth

type UserAuth struct {
	Token    string `mapstructure:"token" form:"token" json:"token"`
	UserID   string `mapstructure:"user_id" form:"user_id" json:"user_id"`
	FullName string `mapstructure:"fullname" form:"fullname" json:"fullname"`
	MaxTodo  int    `json:"max_todo" form:"max_todo"`
}
