package userhandler

import (
	"context"

	"github.com/phathdt/libs/go-sdk/sdkcm"
	"togo/modules/user/usermodel"
)

type UpdateLimitRepo interface {
	UpdateUser(ctx context.Context, cond map[string]interface{}, dataUpdate *usermodel.UserLimit) error
}

type updateLimitHdl struct {
	repo      UpdateLimitRepo
	requester *sdkcm.SimpleUser
}

func NewUpdateLimitHdl(repo UpdateLimitRepo, requester *sdkcm.SimpleUser) *updateLimitHdl {
	return &updateLimitHdl{repo: repo, requester: requester}
}

func (h *updateLimitHdl) Response(ctx context.Context, dataUpdate *usermodel.UserLimit) error {
	err := h.repo.UpdateUser(ctx, map[string]interface{}{"id": h.requester.ID}, dataUpdate)
	if err != nil {
		return sdkcm.ErrCannotUpdateEntity("user", err)
	}

	return nil
}
