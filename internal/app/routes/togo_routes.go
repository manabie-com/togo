package routes

import (
	"net/http"

	"github.com/manabie-com/togo/internal/app/models"
	"github.com/manabie-com/togo/internal/services/transport"
)

type TogoRoutes interface {
	Routes() []*Route
}

type togoRoutesImpl struct {
	transport transport.Transport
}

func NewAppRoutes(transport transport.Transport) *togoRoutesImpl {
	return &togoRoutesImpl{transport}
}

func (r *togoRoutesImpl) Routes() []*Route {
	return []*Route{
		{
			Path:         "/login",
			Method:       http.MethodPost,
			Handler:      r.transport.Login,
			AuthRequired: false,
		},
		{
			Path:         "/tasks",
			Method:       http.MethodGet,
			Handler:      r.transport.ListTasks,
			AuthRequired: true,
		},
		{
			Path:         "/tasks",
			Method:       http.MethodPost,
			Handler:      r.transport.AddTask,
			AuthRequired: true,
		},
		// for test
		{
			Path:         "/refresh-table",
			Method:       http.MethodGet,
			Handler:      models.RefreshTable,
			AuthRequired: false,
		},
	}
}
