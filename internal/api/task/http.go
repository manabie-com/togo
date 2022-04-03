package task

import (
	"net/http"
	"strconv"

	"github.com/TrinhTrungDung/togo/internal/model"
	"github.com/TrinhTrungDung/togo/pkg/server"
	"github.com/labstack/echo/v4"
)

// Auth represents auth interface
type Auth interface {
	User(echo.Context) *model.AuthUser
}

// Service represents task interface
type Service interface {
	List(*model.AuthUser) ([]*model.Task, error)
	View(*model.AuthUser, int) (*model.Task, error)
	Create(*model.AuthUser, TaskData) (*model.Task, error)
	Update(*model.AuthUser, int, TaskData) (*model.Task, error)
	Delete(*model.AuthUser, int) error
}

// HTTP represents task http service
type HTTP struct {
	svc  Service
	auth Auth
}

type TaskData struct {
	Content string `json:"content"`
}

// TaskList contains list of tasks of the current user
type TaskList struct {
	Data []*model.Task `json:"data"`
}

// NewHTTP creates new task http service
func NewHTTP(svc Service, auth Auth, eg *echo.Group) {
	h := HTTP{svc, auth}

	eg.GET("", h.list)
	eg.GET("/:id", h.view)
	eg.POST("", h.create)
	eg.PUT("/:id", h.update)
	eg.DELETE("/:id", h.delete)
}

func (h *HTTP) list(c echo.Context) error {
	resp, err := h.svc.List(h.auth.User(c))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, TaskList{resp})
}

func (h *HTTP) view(c echo.Context) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}

	resp, err := h.svc.View(h.auth.User(c), id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) create(c echo.Context) error {
	r := TaskData{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	resp, err := h.svc.Create(h.auth.User(c), r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) update(c echo.Context) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}

	r := TaskData{}
	if err := c.Bind(&r); err != nil {
		return err
	}

	resp, err := h.svc.Update(h.auth.User(c), id, r)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *HTTP) delete(c echo.Context) error {
	id, err := parseID(c)
	if err != nil {
		return err
	}

	if err := h.svc.Delete(h.auth.User(c), id); err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func parseID(c echo.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, server.NewHTTPValidationError("Invalid ID")
	}

	return id, nil
}
