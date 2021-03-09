package user

type User struct {
	ID       string `gorm:"column:id"`
	Password string `gorm:"column:password"`
	MaxTodo  int    `gorm:"column:max_todo"`
}
