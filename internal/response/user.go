package response

type UserResponse struct {
	ID         int    ` json:"id"`
	Name       string `json:"name"`
	LimitCount int    `json:"limit_count"`
}
