package response

import "time"

type TaskResponse struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	Description string    `json:"description"`
	EndedAt     time.Time `json:"ended_at"`
}
