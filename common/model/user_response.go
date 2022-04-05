package common

import "time"

type UserResponse struct {
	Username         string    `json:"username"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	DailyCap         int64     `json:"daily_cap"`
	DailyQuantity    int64     `json:"daily_quantity"`
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreatedAt        time.Time `json:"created_at"`
}
