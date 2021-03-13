package response

import "github.com/go-kit/kit/endpoint"

var (
	_ endpoint.Failer = &SignUp{}
	_ endpoint.Failer = &SignIn{}
)

type SignUpPayload struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	MaxTodo     int64  `json:"max_todo"`
	CreatedAt   string `json:"created_at"`
	AccessToken string `json:"access_token"`
}

type SignUp struct {
	Data *SignUpPayload `json:"data"`
	Err  error          `json:"-"`
}

func (r SignUp) Failed() error {
	return r.Err
}

type SignInPayload struct {
	AccessToken string `json:"access_token"`
}

type SignIn struct {
	Data *SignInPayload `json:"data"`
	Err  error          `json:"-"`
}

func (r SignIn) Failed() error {
	return r.Err
}
