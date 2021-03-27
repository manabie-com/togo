package domains

import "context"

type (
	LoginRequest struct {
		Username string
		Password string
	}

	User struct {
		Id       int64
		Username string
		MaxTodo  int64
	}

	UserRepository interface {
		VerifyUser(ctx context.Context, request *LoginRequest) (*User, error)
		GetUserById(ctx context.Context, userId int64) (*User, error)
	}
)
