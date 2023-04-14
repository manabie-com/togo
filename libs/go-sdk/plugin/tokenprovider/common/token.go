package common

type TokenPayload struct {
	UId int `json:"user_id"`
}

func (p TokenPayload) UserId() int {
	return p.UId
}
