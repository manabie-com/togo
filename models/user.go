package models

type User struct {
	BaseModelID
	Name       string `json:"name" gorm:"index; not null"`
	LimitCount int    `json:"limit_count" gorm:"not null"`
	BaseModelTime
}

func (User) TableName() string {
	return "public.users"
}
