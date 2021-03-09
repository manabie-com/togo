package mock

import "togo/src"

type ContextMock struct {
	GetTokenDataFunc func() *src.TokenData
}

func (this *ContextMock) GetTokenData() *src.TokenData {
	return this.GetTokenDataFunc()
}

func (this *ContextMock) CheckPermission(scopes []string) error {
	return nil
}

func (this *ContextMock) LoadContext(data interface{}) error {
	return nil
}

func New_ContextMock_With_Valid_UserID() *ContextMock {
	return &ContextMock{
		GetTokenDataFunc: func() *src.TokenData {
			return &src.TokenData{
				UserId: "firstUser",
			}
		},
	}
}

func New_ContextMock_With_InValid_UserID() *ContextMock {
	return &ContextMock{
		GetTokenDataFunc: func() *src.TokenData {
			return &src.TokenData{
				UserId: "",
			}
		},
	}
}
