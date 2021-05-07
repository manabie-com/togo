package entity

type Task struct {
	ID          string ` json:"id,omitempty" gorm:"primary_key"`
	Content     string `json:"content,omitempty"`
	UserID      string `json:"user_id,omitempty"`
	CreatedDate string `json:"created_date,omitempty"`
}

// Save updates the existing or inserts a new row.
func (t *Task) Save() error {
	return Db().Save(t).Error
}

// Create inserts a new row to the database.
func (t *Task) Create() error {
	return Db().Create(t).Error
}

// Updates multiple columns in the database.
func (t *Task) Updates(values interface{}) error {
	return Db().Unscoped().Model(t).UpdateColumns(values).Error
}

// Delete deletes the entity from the database.
func (t *Task) Delete() error {
	return Db().Delete(t).Error
}
