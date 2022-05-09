package entity

type User struct {
	ID              uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserName        string `gorm:"size:50;not null;" json:"username"`
	Password        string `gorm:"size:50;not null;" json:"password"`
	TaskLimitPerDay int64  `gorm:"default:0" json:"task_limit_per_day"`
}
