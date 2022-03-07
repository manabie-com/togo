package routers

import (
	"github.com/gin-gonic/gin"

	"github.com/khoale193/togo/modules/member"
	"github.com/khoale193/togo/modules/task"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()

	api := r.Group("/api")

	// Router Member
	member.MemberRouter(api)

	// Router Member
	task.TaskRouter(api)

	return r
}
