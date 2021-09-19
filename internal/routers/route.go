package routers

import (
	"context"
	"errors"
	"net/http"

	"github.com/manabie-com/togo/internal/repositories"
	"github.com/manabie-com/togo/internal/routers/api"
	endpoints "github.com/manabie-com/togo/internal/routers/endpoints"
	"github.com/manabie-com/togo/internal/services"
	httpPkg "github.com/manabie-com/togo/pkg/http"
	"github.com/manabie-com/togo/pkg/txmanager"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type router struct {
	endpoints endpoints.Endpoints
}

// InitRouter initialize routing information
func InitRouter(db *gorm.DB) http.Handler {
	// Initialize repository
	repo := repositories.InitRepositoryFactory(db)

	tx := txmanager.NewTransactionManager(db)
	// Initialize service
	services := services.InitServiceFactory(repo, tx)
	// Initialize endpoints
	endpoint := endpoints.Endpoints{
		User: endpoints.UserEndpoint{
			GetAuthToken: endpoints.MakeGetAuthTokenEndpoint(services),
		},
		Task: endpoints.TaskEndpoint{
			ListTasks: endpoints.MakeListTasksEndpoint(services),
			AddTask:   endpoints.MakeAddTaskEndpoint(services),
		},
	}

	return &router{endpoint}
}

func (r *router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	logrus.Info(req.Method, req.URL.Path)
	resp.Header().Set("Access-Control-Allow-Origin", "*")
	resp.Header().Set("Access-Control-Allow-Headers", "*")
	resp.Header().Set("Access-Control-Allow-Methods", "*")

	if req.Method == http.MethodOptions {
		resp.WriteHeader(http.StatusOK)
		return
	}

	var (
		ctx     = context.Background()
		resData interface{}
		err     error
	)
	ctx = httpPkg.VerifyTokenAccessContext(ctx, req)
	prePath, behindPath := r.GetGroupPath(req.URL.Path)

	switch prePath {
	case "login":
		resData, err = api.UserRouter(ctx, req, behindPath, &r.endpoints)
	case "tasks":
		resData, err = api.TaskRouter(ctx, req, behindPath, &r.endpoints)
	default:
		err = errors.New("not found")
	}

	httpPkg.EncodeJSONResponse(ctx, resp, resData, err)
}
func (r *router) GetGroupPath(path string) (string, string) {
	if len(path) <= 1 {
		return "", ""
	}

	if string(path[0]) == "/" {
		path = path[1:]
	}

	if string(path[len(path)-1]) == "/" {
		path = path[:len(path)-1]
	}
	i := 0
	for idx, val := range path {
		if idx > 0 && i == 0 && string(val) == "/" {
			i = idx
			break
		}

	}

	if i == 0 {
		return path, ""
	}

	return path[:i], path[i+1:]
}
