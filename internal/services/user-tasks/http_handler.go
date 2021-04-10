package user_tasks

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Handler(
	authorizedV1 *gin.RouterGroup,
	userTaskSrv Service,
) {
	authorizedV1.POST("/create", createTaskHandler(userTaskSrv))
}

func createTaskHandler(userTaskSrv Service) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		r := createTaskRequest{}
		if err := ctx.ShouldBindJSON(&r); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "ERROR",
				"error": map[string]string{
					"code":    "INVALID_REQUEST",
					"message": err.Error(),
				},
			})
			return
		}

		err := userTaskSrv.CreateTask(ctx, r.UserID, r.Content)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"status": "ERROR",
				"error": map[string]string{
					"code":    "INVALID_REQUEST",
					"message": err.Error(),
				},
			})
			return
		}

		ctx.JSON(http.StatusOK, gin.H{
			"status": "OK",
		})
	}
}
