package store

type TodoTask struct {
	ID        uint64 `json:"id"`
	UserID    uint64 `json:"user_id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
}
