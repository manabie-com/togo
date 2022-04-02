package dto

type UsernamePasswordCredential struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type TokenCredential struct {
	Token string `header:"Authorization" binding:"required"`
}

type UserCredential struct {
	ID           int64 `json:"id"`
	MaxDailyTask int   `json:"max_daily_task"`
}
