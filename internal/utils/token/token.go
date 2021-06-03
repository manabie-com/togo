package token

import (
	"net/http"

	"github.com/manabie-com/togo/internal/consts"
)

func GetToken(req *http.Request) string {
	return req.Header.Get(consts.DefaultAuthHeader)

}

//req = req.WithContext(context.WithValue(req.Context(), userAuthKey(0), id))
