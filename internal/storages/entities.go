package storages

import "database/sql"

// Task reflects tasks in DB
type Task struct {
	ID          sql.NullString
	Content     sql.NullString
	UserID      sql.NullString
	CreatedDate sql.NullString
}

// User reflects users data from DB
type User struct {
	ID       sql.NullString
	Password sql.NullString
	MaxTodo  uint32
}
