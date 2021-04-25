package http

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/manabie-com/togo/internal/domain"
)

type APIConf struct {
	JWTSecret string
	TokenExp  time.Duration
}

type HttpAPI struct {
	conf        APIConf
	taskUseCase domain.TaskUseCase
	userUseCase domain.AuthUseCase
}

func BindAPI(conf APIConf, e *echo.Echo, taskUseCase domain.TaskUseCase, userUseCase domain.AuthUseCase) *HttpAPI {
	result := &HttpAPI{
		conf:        conf,
		taskUseCase: taskUseCase,
		userUseCase: userUseCase,
	}
	e.POST("/login", result.Login)
	jwtmw := jwtAuthMiddleware([]byte(conf.JWTSecret))
	e.GET("/tasks", result.ListTasks, jwtmw)
	e.POST("/tasks", result.AddTask, jwtmw)
	return result
}

type LoginInput struct {
	UserID   string `json:"user_id"`
	Password string `json:"password"`
}

func (h *HttpAPI) Login(ctx echo.Context) error {
	var input LoginInput
	err := ctx.Bind(&input)
	if err != nil {
		return fmt.Errorf("error binding value: %s", err)
	}

	ok, err := h.userUseCase.ValidateUser(input.UserID, input.Password)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	if !ok {
		return ctx.JSON(http.StatusUnauthorized, map[string]interface{}{
			"error": "incorrect user_id/pwd",
		})
	}
	token, err := h.createToken(input.UserID)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, map[string]string{
		"data": token,
	})
}

func (h *HttpAPI) createToken(id string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["user_id"] = id
	atClaims["exp"] = time.Now().Add(h.conf.TokenExp).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(h.conf.JWTSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}

type AddTaskInput struct {
	Content string `json:"content"`
}

func (h *HttpAPI) AddTask(ctx echo.Context) error {
	input := AddTaskInput{}
	err := ctx.Bind(&input)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	userID, ok := userIDFromCtx(ctx)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	t := domain.Task{
		ID:          uuid.New().String(),
		UserID:      userID,
		Content:     input.Content,
		CreatedDate: time.Now().Format(domain.DateFormat),
	}
	err = h.taskUseCase.AddTask(t)
	if err != nil {
		if errors.Is(err, domain.TaskLimitReached) {
			return ctx.JSON(http.StatusConflict, map[string]string{
				"error": err.Error(),
			})
		}
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": t,
	})
}

var userAuthKey = "user_id"

func userIDFromCtx(ctx echo.Context) (string, bool) {
	userID, ok := ctx.Get(userAuthKey).(string)
	return userID, ok
}

func (h *HttpAPI) ListTasks(ctx echo.Context) error {
	createdDate := ctx.QueryParam("created_date")
	_, err := time.Parse(domain.DateFormat, createdDate)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, nil)
	}
	userID, ok := userIDFromCtx(ctx)
	if !ok {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	tasks, err := h.taskUseCase.GetTasksByUserIDAndDate(userID, createdDate)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, nil)
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"data": tasks,
	})
}
