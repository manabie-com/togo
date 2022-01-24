package model

type UserConfig struct {
	UserId  uint `json:"user_id" gorm:"user_id"`
	MaxTask uint `json:"max_task" gorm:"max_task"`
}

func (UserConfig) TableName() string {
	return "user_configs"
}
