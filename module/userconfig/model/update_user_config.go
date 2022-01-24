package model

type UpdateUserConfig struct {
	MaxTask *uint `json:"max_task" gorm:"max_task"`
}

func (UpdateUserConfig) TableName() string {
	return UserConfig{}.TableName()
}
