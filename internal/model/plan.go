package model

// Plan represents the plan model
type Plan struct {
	Base
	Name     string `gorm:"varchar(30)"`
	MaxTasks int
}
