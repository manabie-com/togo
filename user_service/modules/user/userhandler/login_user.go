package userhandler

import (
	"context"

	"github.com/phathdt/libs/go-sdk/plugin/tokenprovider"
	"github.com/phathdt/libs/go-sdk/plugin/tokenprovider/jwt"
	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/common"
	"togo/modules/user/usermodel"
)

type LoginUserRepo interface {
	FindUser(ctx context.Context, conditions map[string]interface{}) (*usermodel.User, error)
}

type loginHdl struct {
	repo          LoginUserRepo
	tokenProvider tokenprovider.Provider
	hasher        Hasher
	expiry        int
}

func NewLoginHdl(repo LoginUserRepo, tokenProvider tokenprovider.Provider, hasher Hasher, expiry int) *loginHdl {
	return &loginHdl{
		repo:          repo,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (h *loginHdl) Response(ctx context.Context, data *usermodel.UserLogin) (tokenprovider.Token, error) {
	user, err := h.repo.FindUser(ctx, map[string]interface{}{"email": data.Email})

	if err != nil {
		return nil, common.ErrEmailOrPasswordInvalid
	}

	passHashed := h.hasher.Hash(data.Password + user.Salt)

	if user.Password != passHashed {
		return nil, common.ErrEmailOrPasswordInvalid
	}

	payload := &jwt.TokenPayload{
		UId: user.ID,
	}

	accessToken, err := h.tokenProvider.Generate(payload, h.expiry)
	if err != nil {
		return nil, sdkcm.ErrInternal(err)
	}

	return accessToken, nil
}
