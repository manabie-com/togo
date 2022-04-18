package model

type NewUser struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	Password string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	TaskLimit int `json:"task_limit"`
}

type UserDetails struct {
	Username string `json:"username"`
	Name string `json:"name"`
	Email string `json:"email"`
	TaskLimit int `json:"task_limit"`
}

type UserNewPassword struct {
	CurrentPassword string `json:"current_password"`
	NewPassword string `json:"new_password"`
	RepeatPassword string `json:"repeat_password"` 
}

type UserLogin struct {
	Username string `json:"username"`
	Password string `json:"password"` 
}