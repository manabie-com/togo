package model

// User model that will be used when a new user is "registered"
type NewUser struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	TaskLimit int `json:"task_limit"`
}

// User model that will be used when fetching the details of a user, ommitting unneeded fields
type UserDetails struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	TaskLimit int `json:"task_limit"`
}

// User model that will be used on user password change
type UserNewPassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword string `json:"new_password"`
	RepeatPassword string `json:"repeat_password"` 
}

// User model that will be used on login
type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"` 
}