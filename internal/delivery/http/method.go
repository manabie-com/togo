package http

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/manabie-com/togo/pkg/model"

	"github.com/gin-gonic/gin"
)

func (h *Handler) Login(ctx *gin.Context) {
	var user model.User
	rawData, _ := ctx.GetRawData()
	if err := json.Unmarshal(rawData, &user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"massage": "invalid authenticate data",
		})
		return
	}
	token, err := h.userService.Authentication(&user)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"massage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "authenticated",
		"token":   token,
	})
}

func (h *Handler) ListTasks(ctx *gin.Context) {
	userI, _ := ctx.Get("user")
	userId := userI.(*model.User).ID
	var createdDate *time.Time

	createdDateString := ctx.Query("created_date")
	if t, err := time.Parse("2006-01-02T15:04:05", createdDateString+"T00:00:00"); err == nil {
		createdDate = &t
	}

	tasks, err := h.TaskService.FindTask(userId, createdDate)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"massage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"massage": "success",
		"data":    tasks,
	})
}

func (h *Handler) CreateTasks(ctx *gin.Context) {
	userI, _ := ctx.Get("user")
	user := userI.(*model.User)

	rawData, _ := ctx.GetRawData()
	var body struct {
		Content string `json:"content"`
	}
	if err := json.Unmarshal(rawData, &body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"massage": err.Error(),
		})
		return
	}

	err := h.TaskService.CreateTask(user, body.Content)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"massage": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"massage": "success",
	})
}

func (h *Handler) Authorise(ctx *gin.Context) {
	token := ctx.GetHeader("Authorization")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"massage": "require Authorization header",
		})
		ctx.Abort()
	}

	// get only token , remove Bearer
	token = strings.Replace(token, "Bearer ", "", 1)

	// validate token
	user, err := h.userService.Authorise(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"code":    http.StatusUnauthorized,
			"massage": err.Error(),
		})
		ctx.Abort()
	}

	ctx.Set("user", user)
}
