package models

type User struct {
	ID            uint32  `json:"id"`
	Email         string  `json:"email" validate:"required"`
	Name          string  `json:"name" validate:"required"`
	Password      string  `json:"password" validate:"required"`
	IsPayment     bool    `json:"isPayment"`
	LimitDayTasks uint    `json:"limitDayTasks"`
	Tasks         *[]Task `json:"tasks"`
}

func (u *User) validate() error {
	// v := validator.New()
	// err := v.Struct(u)
	// for _, e := range err.(validator.ValidationErrors) {
	// 	fmt.Print(e)
	// }
	return nil
}
