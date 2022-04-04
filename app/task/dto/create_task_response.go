package dto

type CreateTaskResponse struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	OwnerID   int64  `json:"owner_id"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
