package services

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jericogantuangco/togo/internal/token"
	"github.com/stretchr/testify/require"
)

func TestSAuthMiddlewareOK(t *testing.T) {
	server, err := NewServer(nil)
	require.NoError(t, err)

	authPath := "/auth"
	server.Router.GET(
		authPath,
		authMiddleware(server.TokenMaker),
		func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{})
		},
	)

	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodGet, authPath, nil)
	require.NoError(t, err)

	addAuth(t, request, server.TokenMaker, authorizationTypeBearer, "testUser", 10*time.Minute)
	server.Router.ServeHTTP(recorder, request)
	checkResponse(t, recorder)
}

func addAuth(t *testing.T, req *http.Request, tokenMaker token.Maker, authType string, username string, duration time.Duration) {
	token, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authHeader := fmt.Sprintf("%s %s", authType, token)
	req.Header.Set(authorizationHeaderKey, authHeader)
}

func checkResponse(t *testing.T, recorder *httptest.ResponseRecorder) {
	require.Equal(t, http.StatusOK, recorder.Code)
}
