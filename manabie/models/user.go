package models

// User struct
type User struct {
	ID       string `gorm:"column:id;primary_key;not null" json:"id"`
	Password string `gorm:"column:password" json:"password"`
	MaxToDo  int    `gorm:"column:max_todo" json:"max_todo"`
}

// TableName func
func (User) TableName() string {
	return "users"
}
