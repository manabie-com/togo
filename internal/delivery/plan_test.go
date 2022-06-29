package delivery

import (
	"bytes"
	"lntvan166/togo/internal/config"
	"lntvan166/togo/pkg/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestGetPlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskUsecase := mock.NewMockTaskUsecase(ctrl)
	userUsecase.EXPECT().GetPlan(user.Username).Return(user.Plan, nil)

	userDelivery := NewUserDelivery(userUsecase, taskUsecase)

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/plan", nil)
	r.Header.Set("Authorization", "Bearer "+token)
	context.Set(r, "username", user.Username)

	userDelivery.GetPlan(w, r)

	res := w.Result()

	assert.Equal(t, http.StatusOK, w.Code)

	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(res.Body)
	body := strings.TrimSpace(bodyBuffer.String())

	assert.Equal(t, `"plan: `+user.Plan+`"`, body)
}

func TestUpGradePlan(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userUsecase := mock.NewMockUserUsecase(ctrl)
	taskUsecase := mock.NewMockTaskUsecase(ctrl)
	userUsecase.EXPECT().UpgradePlan(user.ID, userAfterUpgrade.Plan, int(userAfterUpgrade.MaxTodo)).
		Return(nil).AnyTimes()

	userDelivery := NewUserDelivery(userUsecase, taskUsecase)
	config.ADMIN = "admin"
	config.VIP_LIMIT = 20

	w := httptest.NewRecorder()
	r, _ := http.NewRequest("POST", "/plan", nil)

	vars := map[string]string{
		"id": "1",
	}

	r = mux.SetURLVars(r, vars)
	r.Header.Set("Authorization", "Bearer "+token)
	context.Set(r, "username", user.Username)

	userDelivery.UpgradePlan(w, r)

	assert.Equal(t, http.StatusOK, w.Code)
}
