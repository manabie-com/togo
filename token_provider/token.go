package tokenprovider

import (
	"time"
)

type Token struct {
	Token   string    `json:"token"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type JwtPayload struct {
	UserId int `json:"userId,omitempty"`
}

func (j JwtPayload) GetUserId() int {
	return j.UserId
}

type IPayload interface {
	GetUserId() int
}
