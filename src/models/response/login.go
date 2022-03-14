package response

type LoginResp struct {
	UserID int    `json:"user_id"`
	Token  string `json:"token"`
}
