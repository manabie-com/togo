package plan

import (
	"net/http"

	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/labstack/echo/v4"
)

// Service represents plan application interface
type Service interface {
	List() ([]*model.Plan, error)
}

// HTTP represents plan http service
type HTTP struct {
	svc Service
}

// PlanList contains list of current provided plans
type PlanList struct {
	Data []*model.Plan `json:"data"`
}

// NewHTTP creates new plan http service
func NewHTTP(svc Service, eg *echo.Group) {
	h := HTTP{svc}

	eg.GET("", h.list)
}

func (h *HTTP) list(c echo.Context) error {
	resp, err := h.svc.List()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, PlanList{resp})
}
