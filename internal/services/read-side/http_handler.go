package read_side

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

func Handler(
	authorizedV1 *gin.RouterGroup,
	readRepo ReadRepo,
) {
	authorizedV1.GET("/task-list", getTaskList(readRepo))
}

func getTaskList(readRepo ReadRepo) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		values := ctx.Request.URL.Query()
		list, cursor, err := readRepo.GetTaskList(ctx, values)
		if err != nil {
			logrus.WithError(err).Errorf("GetTaskList, values: %v", values)
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
			"data":   list,
			"meta_data": map[string]string{
				"cursor": cursor,
			},
		})
	}
}
