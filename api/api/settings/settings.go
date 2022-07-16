package settings

import (
	"net/http"
	"strconv"

	"manabie/todo/models"
	"manabie/todo/pkg/utils"
	"manabie/todo/service/setting"

	"github.com/labstack/echo/v4"
)

type handler struct {
	Setting setting.SettingService
}

func NewSettingHandler(e *echo.Echo, st setting.SettingService) {
	h := &handler{
		Setting: st,
	}
	e.GET("/users/:id/settings", h.Show)
	e.POST("/users/:id/settings", h.Create)
	e.PUT("/settings/:id", h.Update)
}

func (h *handler) Show(c echo.Context) error {
	mid := c.Param("id")

	memberID, err := strconv.Atoi(mid)
	if err != nil {
		return utils.ResponseFailure(c, http.StatusBadRequest, err)
	}

	st, err := h.Setting.Show(c.Request().Context(), memberID)
	if err != nil {
		return utils.ResponseFailure(c, http.StatusInternalServerError, err)
	}

	return utils.ResponseSuccess(c, st)
}

func (h *handler) Create(c echo.Context) error {
	mid := c.Param("id")
	st := new(models.SettingCreateRequest)

	memberID, err := strconv.Atoi(mid)
	if err != nil {
		return utils.ResponseFailure(c, http.StatusBadRequest, err)
	}

	if err := c.Bind(st); err != nil {
		return utils.ResponseFailure(c, http.StatusBadRequest, err)
	}

	if err := h.Setting.Create(c.Request().Context(), memberID, st); err != nil {
		return utils.ResponseFailure(c, http.StatusInternalServerError, err)
	}

	return utils.ResponseSuccess(c, models.StatusResponse{
		Status: "ok",
	})
}

func (h *handler) Update(c echo.Context) error {
	id := c.Param("id")
	st := new(models.SettingUpdateRequest)

	settingID, err := strconv.Atoi(id)
	if err != nil {
		return utils.ResponseFailure(c, http.StatusBadRequest, err)
	}

	if err := c.Bind(st); err != nil {
		return utils.ResponseFailure(c, http.StatusBadRequest, err)
	}

	if err := h.Setting.Update(c.Request().Context(), settingID, st); err != nil {
		return utils.ResponseFailure(c, http.StatusInternalServerError, err)
	}

	return utils.ResponseSuccess(c, models.StatusResponse{
		Status: "ok",
	})
}
