package model

// Task reflects tasks in DB
type Task struct {
	ID          string    `json:"id" db:"id"`
	Content     string    `json:"content" db:"content"`
	UserID      string    `json:"user_id" db:"user_id"`
	CreatedDate string `json:"created_date" db:"created_date"`
}

func (t Task) IsValid() error {
	if len(t.ID) == 0 {
		return NewError(ErrInvalidUserModel, "invalid field: id")
	}

	if len(t.Content) == 0 {
		return NewError(ErrInvalidUserModel, "empty field: content")
	}

	if len(t.UserID) == 0 {
		return NewError(ErrInvalidUserModel, "invalid field: user_id")
	}

	return nil
}

// User reflects users data from DB
type User struct {
	ID           string `json:"id" db:"id"`
	PasswordHash string `json:"-" db:"password_hash"`
	MaxTodo      int    `json:"max_todo" db:"max_todo"`
}

func (u User) IsValid() error {
	if len(u.ID) == 0 {
		return NewError(ErrInvalidUserModel, "invalid field: id")
	}

	if len(u.PasswordHash) == 0 {
		return NewError(ErrInvalidUserModel, "invalid field: password")
	}

	return nil
}
