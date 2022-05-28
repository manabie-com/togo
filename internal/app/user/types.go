package user

import (
	"time"

	"github.com/dinhquockhanh/togo/internal/pkg/token"
	"github.com/dinhquockhanh/togo/internal/pkg/uuid"
)

type (
	CreateUserReq struct {
		UserName string `json:"username" binding:"required,min=3"`
		FullName string `json:"full_name" binding:"required,min=2"`
		Password string `json:"password" binding:"required,min=5"`
		Email    string `json:"email" binding:"required,email"`
	}

	UpdateUserTierReq struct {
		TierID   int32  `json:"tier_id" binding:"required,number"`
		UserName string `json:"username" binding:"required,min=6"`
	}

	GetUserByUserNameReq struct {
		UserName string `uri:"username" binding:"required,min=3"`
	}

	ListUsersReq struct {
		PageNumber int `uri:"page_number" binding:"required,min=1"`
		PageSize   int `uri:"page_size" binding:"required,min=1"`
	}

	DeleteUserByNameReq struct {
		UserName string `uri:"username" binding:"required,min=6"`
	}

	UserSafe struct {
		Username  string    `json:"username"`
		FullName  string    `json:"full_name"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"created_at"`
		TierID    int32     `json:"tier_id"`
	}

	User struct {
		Username       string
		FullName       string
		HashedPassword string
		Email          string
		CreatedAt      time.Time
		TierID         int32
	}
)

func (u *User) Safe() *UserSafe {
	return &UserSafe{
		Username:  u.Username,
		FullName:  u.FullName,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
		TierID:    u.TierID,
	}
}

func (u *User) ToPayload(duration time.Duration) *token.Payload {
	return &token.Payload{
		ID:        uuid.New(),
		Username:  u.Username,
		TierID:    int(u.TierID),
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(duration),
	}
}

// PlayLoadToUser extract information from a  token Playload and return a user
func PlayLoadToUser(payload *token.Payload) *User {
	return &User{
		Username: payload.Username,
		TierID:   int32(payload.TierID),
	}
}
