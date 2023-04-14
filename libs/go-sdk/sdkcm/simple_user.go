package sdkcm

type SimpleUser struct {
	SQLModel `json:",inline"`
	Email    string `json:"email" gorm:"column:email;"`
}

func (u SimpleUser) TableName() string {
	return "users"
}
