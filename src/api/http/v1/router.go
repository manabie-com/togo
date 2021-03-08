package v1

import (
	"togo/src/controller/task"
	"togo/src/controller/user"
	"togo/src/schema"

	"net/http"

	gErrors "togo/src/infra/error"
	"togo/src/infra/service"

	"github.com/labstack/echo/v4"
)

type RouterV1 struct {
	framework      *echo.Echo
	taskController task.ITaskController
	userController user.IUserController
}

func (this *RouterV1) LoadAPI() {
	group := this.framework.Group("/api")

	group.POST("/login", func(c echo.Context) error {
		request := new(schema.LoginRequest)

		if err := c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		data, err := this.userController.Login(request)
		if err != nil {
			switch err.(type) {
			case *gErrors.UnauthorizedError:
				return c.JSON(http.StatusUnauthorized, err.Error())
			case *gErrors.NotFoundError:
				return c.JSON(http.StatusNotFound, err.Error())
			default:
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusOK, data)
	})

	group.POST("/tasks", func(c echo.Context) error {

		context := service.NewServiceContext()
		if err := context.LoadContext(c.Request().Header); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		request := new(schema.AddTaskRequest)
		if err := c.Bind(request); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		if err := c.Validate(request); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		data, err := this.taskController.AddTaskByOwner(context, request)
		if err != nil {
			switch err.(type) {
			case *gErrors.UnauthorizedError:
				return c.JSON(http.StatusUnauthorized, err.Error())
			case *gErrors.NotFoundError:
				return c.JSON(http.StatusNotFound, err.Error())
			case *gErrors.BadRequestError:
				return c.JSON(http.StatusBadRequest, err.Error())
			default:
				return c.JSON(http.StatusInternalServerError, err.Error())
			}
		}

		return c.JSON(http.StatusOK, data)
	})
}

func NewRouterV1(framework *echo.Echo) *RouterV1 {
	return &RouterV1{
		framework:      framework,
		taskController: task.NewTaskController(),
		userController: user.NewUserController(),
	}
}
