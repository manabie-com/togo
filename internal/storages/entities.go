package storages

import "gorm.io/gorm"

// Task reflects tasks in DB
type Task struct {
	gorm.Model
	ID          string `json:"id"           mapstructure:"id"           gorm:"primary_key:true;column:id;not null"`
	Content     string `json:"content"      mapstructure:"content"      gorm:"column:content;not null"`
	UserID      string `json:"user_id"      mapstructure:"user_id"      gorm:"column:user_id;not null"`
	CreatedDate string `json:"created_date" mapstructure:"created_date" gorm:"column:created_date;not null"`
	User		*User  `gorm:"foreignKey:UserID"`
}

func(Task) TableName() string {
	return "tasks"
}

func(Task) BeforeCreate(tx *gorm.DB) error {
	return nil
}

// User reflects users data from DB
type User struct {
	gorm.Model
	ID       string	`json:"id"       gorm:"primary_key;column:id;not null"  mapstructure:"id"`
	Password string	`json:"password" gorm:"column:password;not null"        mapstructure:"password"`
	MaxTodo  int32 	`json:"max_todo" gorm:"column:max_todo;not null"        mapstructure:"max_todo"`
}

func(User) TableName() string{
	return "users"
}

func(User) BeforeCreate(tx *gorm.DB) error {
	return nil
}
