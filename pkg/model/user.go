package model

// User reflects users data from DB
type User struct {
	ID       int    `json:"id" gorm:"column=id;PRIMARY"`
	UserName string `json:"user_name" gorm:"column=user_name;UNIQUE"`
	Password string `json:"password" gorm:"column=password;"`
	Salt     string `json:"salt" gorm:"column=salt"`
	MaxTodo  int    `json:"max_todo" gorm:"column=max_todo;"`
}

func (m *User) TableName() string {
	return `users`
}

func (m *User) IsEmptyAuthenticationInfo() bool {
	return m.UserName == "" || m.Password == "ß"
}
