package model

// User represents the user model
type User struct {
	Base
	FirstName string `gorm:"type:varchar(255)"`
	LastName  string `gorm:"type:varchar(255)"`
	Email     string `gorm:"type:varchar(254);unique_index;not null"`
	Username  string `gorm:"type:varchar(255);unique_index;not null"`
	Password  string `gorm:"type:varchar(255);not null"`
	Tasks     []*Task
}
