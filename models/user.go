package models

type User struct {
	ID            uint32  `json:"id"`
	Name          string  `json:"name" validate:"required"`
	Password      string  `json:"password" validate:"required"`
	IsPayment     bool    `json:"isPayment" default:"false"`
	LimitDayTasks uint    `json:"limitDayTasks" default:"10"`
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

type user interface {
	/* map[string]interface{} return type like
	{
		status: "success"||"failure",
		message: Optional
		data: {
			data
		}
	}
	*/
	Create(map[string]interface{}) map[string]interface{}
	// update to premium user => increase user daily tasks
	Update(map[string]interface{}) map[string]interface{}

	GetUser(id string) map[string]interface{}
}
