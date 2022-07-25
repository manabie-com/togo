package transformer

import (
	"github.com/huuthuan-nguyen/manabie/app/model"
	"time"
)

type TokenTransformer struct {
	AccessToken string    `json:"access_token"`
	ExpiredAt   time.Time `json:"expired_at"`
}

// Transform /**
func (token *TokenTransformer) Transform(e any) any {
	tokenModel, ok := e.(model.Token)
	if !ok {
		return e
	}

	token.AccessToken = tokenModel.AccessToken
	token.ExpiredAt = tokenModel.ExpiredAt
	return *token
}
