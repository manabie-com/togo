package subscription

import (
	"net/http"

	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/labstack/echo/v4"
)

// Auth represents auth interface
type Auth interface {
	User(echo.Context) *model.AuthUser
}

// Service represents subscription interface
type Service interface {
	Subscribe(*model.AuthUser, SubscriptionData) error
}

// HTTP represents subscription http service
type HTTP struct {
	svc  Service
	auth Auth
}

// SubscriptionData contains user's plan subscription data from JSON request
type SubscriptionData struct {
	PlanName string `json:"plan_name" validate:"required"`
}

// NewHTTP creates new subscription http service
func NewHTTP(svc Service, auth Auth, eg *echo.Group) {
	h := HTTP{svc, auth}

	eg.POST("", h.subscribe)
}

func (h *HTTP) subscribe(c echo.Context) error {
	body := SubscriptionData{}
	if err := c.Bind(&body); err != nil {
		return err
	}

	if err := h.svc.Subscribe(h.auth.User(c), body); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}
