package request

type UserLoginRequest struct {
	Username string `query:"username" json:"username"`
	Password string `query:"password" json:"password"`
}
