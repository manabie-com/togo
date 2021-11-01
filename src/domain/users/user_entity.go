package users

import "time"

type (
	User struct {
		ID          string    `sql:"primary_key;auto_increment" json:"id"`
		Username    string    `sql:"username;" json:"username"`
		Password    string    `sql:"password" json:"-"`
		MaxTodo     int       `sql:"max_todo" json:"maxTodo"`
		CreatedDate time.Time `sql:"created_date,default:CURRENT_TIMESTAMP" json:"createdDate"`
	}

	Users []User
)
