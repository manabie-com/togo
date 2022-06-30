package entities

type User struct {
	CommonModel
	Username string `json:"username" validate:"required, lte=5,gte=10" gorm:"unique"`
	Password string `json:"password"`
	Limit    int    `json:"limit" validate:"required" gorm:"default:5"`
}
