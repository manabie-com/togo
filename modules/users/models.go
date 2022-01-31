package users

type Users struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	IsActive bool   `json:"isActive" gorm:"default:true"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
