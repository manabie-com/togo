package model

type LoginRequest struct {
	UserName string `json:"username" query:"username"`
	Password string `json:"password" query:"password"`
}
