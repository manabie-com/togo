package usermodel

type UserCreate struct {
	Email    string `form:"username" binding:"email" gorm:"column:email"`
	Password string `form:"password" binding:"required,min=6,max=32" gorm:"column:password"`
	Salt     string `form:"-" gorm:"column:salt;"`
}

func (UserCreate) TableName() string {
	return "users"
}
