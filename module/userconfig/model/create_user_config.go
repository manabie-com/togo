package model

type CreateUserConfig struct {
	UserId  *uint `json:"user_id" gorm:"user_id"`
	MaxTask uint `json:"max_task" gorm:"max_task"`
}

func (CreateUserConfig) TableName() string {
	return UserConfig{}.TableName()
}
