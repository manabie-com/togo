package tokenprovider

import (
	"time"
)

type Token struct {
	Token   string    `json:"tokenprovider"`
	Created time.Time `json:"created"`
	Expiry  int       `json:"expiry"`
}

type JwtPayload struct {
	UserId         int    `json:"userId,omitempty"`
	Role           string `json:"role,omitempty"`
	RefreshTokenId string `json:"refreshTokenId,omitempty"`
}

func (j JwtPayload) GetUserId() int {
	return j.UserId
}

func (j JwtPayload) GetRefreshTokenId() string {
	return j.RefreshTokenId
}

type IPayload interface {
	GetUserId() int
	GetRefreshTokenId() string
}
