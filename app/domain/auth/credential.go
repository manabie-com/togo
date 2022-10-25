package auth

type AuthenticationCredential struct {
	ID           int64 `json:"id"`
	MaxDailyTask int   `json:"max_daily_task"`
}
