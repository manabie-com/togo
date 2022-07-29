package response

type UserResponse struct {
	ID         int    `json:"id"`
	LimitCount int    `json:"limit_count"`
	Name       string `json:"name"`
}
