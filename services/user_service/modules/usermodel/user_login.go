package usermodel

type UserLogin struct {
	Email    string `form:"username" binding:"email" gorm:"column:email"`
	Password string `form:"password" binding:"required,min=6,max=32" gorm:"column:password"`
}

func (UserLogin) TableName() string {
	return User{}.TableName()
}
