package mock

import "togo/src"

type JwtMock struct {
	GetCreateTokenFunc func(data *src.TokenData) (string, error)
}

func (this *JwtMock) CreateToken(data *src.TokenData) (string, error) {
	return this.GetCreateTokenFunc(data)
}

func (this *JwtMock) VerifyToken(token string) (*src.TokenData, error) {
	return nil, nil
}

var (
	TOKEN = "this-is-secret-token"
)

func New_JwtMock_With_CreateTokenOK() *JwtMock {
	return &JwtMock{
		GetCreateTokenFunc: func(data *src.TokenData) (string, error) {
			return TOKEN, nil
		},
	}
}

func New_JwtMock_With_CreateTokenError() *JwtMock {
	return &JwtMock{
		GetCreateTokenFunc: func(data *src.TokenData) (string, error) {
			return "", ERROR_500
		},
	}
}
