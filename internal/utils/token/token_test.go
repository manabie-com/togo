package token

import (
	"net/http"
	"testing"

	"github.com/manabie-com/togo/internal/consts"
	"github.com/manabie-com/togo/internal/utils/random"
	"github.com/stretchr/testify/assert"
)

func TestGetToken(t *testing.T) {
	t.Parallel()
	methodList := []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodDelete}
	for _, method := range methodList {
		t.Run("Should retrieve correct authentication header", func(t *testing.T) {
			request, _ := http.NewRequest(method, "", nil)
			expectedResult := random.RandString(10)
			request.Header.Add(consts.DefaultAuthHeader, expectedResult)
			got := GetToken(request)
			assert.Equal(t, expectedResult, got)
		})
	}
}
