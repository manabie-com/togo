package service

import (
	"errors"
	"togo/src"

	gErrors "togo/src/infra/error"

	"github.com/mitchellh/mapstructure"
)

type Context struct {
	token        string
	tokenData    *src.TokenData
	errorFacotry src.IErrorFactory
	jwtService   src.IJWTService
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func (this *Context) GetTokenData() *src.TokenData {
	return this.tokenData
}

func (this *Context) CheckPermission(scopes []string) error {
	data, err := this.jwtService.VerifyToken(this.token)
	if err != nil {
		return this.errorFacotry.UnauthorizedError(src.TOKEN_INVALID, err)
	}

	for _, scope := range scopes {
		if !contains(data.Permissions, scope) {
			return this.errorFacotry.UnauthorizedError(src.NO_PERMISSION, errors.New(scope))
		}
	}

	this.tokenData = data

	return nil
}

func (this *Context) LoadContext(data interface{}) error {
	header := new(src.HeaderData)

	var tempConvert struct {
		AccessToken []string `mapstructure:"access_token"`
	}

	if err := mapstructure.Decode(data, &tempConvert); err != nil {
		return this.errorFacotry.InternalServerError(src.TOKEN_INVALID, err)
	}

	if len(tempConvert.AccessToken) == 0 {
		return this.errorFacotry.UnauthorizedError(src.TOKEN_NOT_PROVIED, errors.New("token is not provied"))
	}

	header.AccessToken = tempConvert.AccessToken[0]
	this.token = header.AccessToken

	return nil
}

func NewServiceContext() src.IContextService {
	return &Context{
		"",
		&src.TokenData{},
		gErrors.NewErrorFactory(),
		NewJWTService(),
	}
}
