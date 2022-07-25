package transformer

import (
	"github.com/huuthuan-nguyen/manabie/app/model"
	"time"
)

type UserTransformer struct {
	ID         int       `json:"id"`
	Email      string    `json:"email"`
	IsActive   bool      `json:"is_active"`
	DailyLimit int       `json:"daily_limit"`
	CreateAt   time.Time `json:"create_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

// Transform /**
func (user *UserTransformer) Transform(e any) any {
	userModel, ok := e.(model.User)
	if !ok {
		return e
	}

	user.ID = userModel.ID
	user.Email = userModel.Email
	user.IsActive = userModel.IsActive
	user.DailyLimit = userModel.DailyLimit
	user.CreateAt = userModel.CreatedAt
	user.UpdatedAt = userModel.UpdatedAt
	return *user
}
