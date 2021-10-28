package handler

import (
	"errors"
	"net/http"

	"github.com/manabie-com/togo/internal/core/port"

	"github.com/gin-gonic/gin"
)

// ToDoService implement HTTP server
type ToDoService struct {
	jwtService  port.JwtService
	taskService port.TaskService
}

func (p *httpHandler) responseSuccess(c *gin.Context, data interface{}) {
	if data == nil {
		c.JSON(http.StatusOK, map[string]interface{}{})
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"data": data,
	})
}

func (p *httpHandler) responseError(c *gin.Context, statusCode int, err error) {
	c.JSON(statusCode, map[string]string{
		"error": err.Error(),
	})
}

func (p *httpHandler) getUserIdFromRequest(c *gin.Context) (string, error) {
	token := c.GetHeader("Authorization")
	return p.jwtService.ParseToken(token)
}

func (p *httpHandler) login(c *gin.Context) {
	var req reqLogin
	err := c.BindJSON(&req)
	if err != nil {
		p.responseError(c, http.StatusBadRequest, err)
		return
	}

	userId, err := p.taskService.Login(c.Request.Context(), req.Username, req.Password)
	if err != nil {
		p.responseError(c, http.StatusInternalServerError, err)
		return
	}
	if len(userId) == 0 {
		p.responseError(c, http.StatusUnauthorized, errors.New("incorrect username or password"))
		return
	}

	token, err := p.jwtService.CreateToken(userId)
	if err != nil {
		p.responseError(c, http.StatusInternalServerError, err)
		return
	}
	p.responseSuccess(c, token)
}

func (p *httpHandler) getTasks(c *gin.Context) {
	userId, err := p.getUserIdFromRequest(c)
	if err != nil {
		p.responseError(c, http.StatusBadRequest, err)
		return
	}
	tasks, err := p.taskService.RetrieveTasks(c.Request.Context(), userId, c.Query("created_date"))
	if err != nil {
		p.responseError(c, http.StatusInternalServerError, err)
		return
	}
	p.responseSuccess(c, tasks)
}

func (p *httpHandler) addTask(c *gin.Context) {
	var req reqAddTask
	err := c.BindJSON(&req)
	if err != nil {
		p.responseError(c, http.StatusBadRequest, err)
		return
	}

	userId, err := p.getUserIdFromRequest(c)
	if err != nil {
		p.responseError(c, http.StatusBadRequest, err)
		return
	}

	createdTask, err := p.taskService.AddTask(c.Request.Context(), userId, req.Content)
	if err != nil {
		p.responseError(c, http.StatusInternalServerError, err)
		return
	}
	p.responseSuccess(c, createdTask)
}
