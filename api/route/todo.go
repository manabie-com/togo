package route

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/api/handler"
	"github.com/manabie-com/togo/api/middleware"
	"github.com/manabie-com/togo/internal/repo"
	todorepo "github.com/manabie-com/togo/internal/repo/todo"
	"github.com/manabie-com/togo/internal/service"
)

func RegisterTodo(router *gin.RouterGroup, c repo.Conn) {
	router.Use(middleware.CheckUserId())
	router.POST(Todo, handler.AddTodo(&service.DefaultTodo{
		Repo: &todorepo.InmemTodo{
			Conn: c,
		},
	}))
	router.GET(Todo, handler.GetTodo(&service.DefaultTodo{
		Repo: &todorepo.InmemTodo{
			Conn: c,
		},
	}))
}
