package auth

type (
	LoginReq struct {
		UserName string `json:"username" binding:"required,min=3"`
		Password string `json:"password" binding:"required,min=5"`
	}
)
