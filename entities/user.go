package entities

type User struct {
	ID       uint `json:"id"`
	UserName uint `json:"userName"`
}

func (b *User) TableName() string {
	return "users"
}
