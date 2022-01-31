package users

type Users struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	IsActive bool   `json:"isActive" gorm:"default:true"`
	Name     string `json:"name"`
	Limit    uint   `json:"limit" gorm:"default:5"`
	Password string `json:"password"`
}
