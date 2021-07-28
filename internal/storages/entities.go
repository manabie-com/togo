package storages

// Task reflects tasks in DB
type Task struct {
	ID          string `gorm:"column:id" json:"id"`
	Content     string `gorm:"column:content;not null" json:"content"`
	UserID      string `gorm:"column:user_id" json:"user_id"`
	CreatedDate string `gorm:"column:created_date" json:"created_date"`
}

func (Task) TableName() string {
	return "tasks"
}

// User reflects users data from DB
type User struct {
	ID       string `gorm:"column:id"`
	Password string `gorm:"column:password"`
	MaxTodo  int    `gorm:"column:max_todo"`
	Tasks    []Task `gorm:"foreignkey:UserID"  json:"tasks"`
}

func (User) TableName() string {
	return "users"
}
