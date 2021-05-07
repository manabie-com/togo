package entity

const DefaultMaxTodo = 5

type User struct {
	ID       string `json:"id,omitempty" gorm:"primary_key"`
	Password string `json:"password,omitempty"`
	MaxTodo  int    `json:"max_todo,omitempty"`
	//GormCustomTime
}

// Save updates the existing or inserts a new row.
func (u *User) Save() error {
	return Db().Save(u).Error
}

// Create inserts a new row to the database.
func (u *User) Create() error {
	return Db().Create(u).Error
}

// Updates multiple columns in the database.
func (u *User) Updates(values interface{}) error {
	return Db().Unscoped().Model(u).UpdateColumns(values).Error
}

// Delete deletes the entity from the database.
func (u *User) Delete() error {
	return Db().Delete(u).Error
}
