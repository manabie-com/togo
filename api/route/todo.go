package route

import (
	"github.com/gin-gonic/gin"
	"github.com/manabie-com/togo/api/handler"
	"github.com/manabie-com/togo/internal/repo"
	todorepo "github.com/manabie-com/togo/internal/repo/todo"
	"github.com/manabie-com/togo/internal/service"
)

func RegisterTodo(router *gin.RouterGroup, c repo.Conn) {
	// Handler godoc
	// @Summary endpoint to create todo item
	// @Description API to add an item to todo list
	// @ID create-todo
	// @Param Authorization header string true "Authorization"
	// @Param body body model.Todo true "todo information"
	// @Produce json
	// @Success 200 {object} model.Todo
	// @Router /v1/todo [POST]
	router.POST(todo, handler.AddTodo(&service.DefaultTodo{
		Repo: &todorepo.InmemTodo{
			Conn: c,
		},
	}))
	// Handler godoc
	// @Summary endpoint to create todo item
	// @Description API to add an item to todo list
	// @ID create-todo
	// @Param Authorization header string true "Authorization"
	// @Param body body model.Todo true "todo information"
	// @Produce json
	// @Success 200 {object} model.Todo
	// @Router /v1/todo [GET]
	router.GET(todo, handler.GetTodo(&service.DefaultTodo{
		Repo: &todorepo.InmemTodo{
			Conn: c,
		},
	}))
}
