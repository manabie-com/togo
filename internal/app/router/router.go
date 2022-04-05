package router

import (
	"net/http"
	handlers "togo/internal/pkg/deliveries/http"
	"togo/internal/pkg/domain/dtos"
	"togo/internal/pkg/repositories"
	"togo/internal/pkg/usecases"
	"togo/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	Engine *gin.Engine
	DB     *gorm.DB
}

func (r *Router) InitRoute() {
	r.Engine.Use(gin.Logger())
	r.Engine.Use(gin.Recovery())

	// auth
	ur := repositories.NewUserRepository(r.DB)
	au := usecases.NewAuthUsecase(ur)
	ah := handlers.NewAuthHandler(au)

	tr := repositories.NewToDoRepository(r.DB)
	tu := usecases.NewTodoUsecase(tr)
	th := handlers.NewTodoHandler(tu)

	r.Engine.GET("/health-check", func(c *gin.Context) {
		data := dtos.BaseResponse{
			Status: http.StatusOK,
			Data:   gin.H{"message": "Health check OK!"},
			Error:  nil,
		}
		c.JSON(http.StatusOK, data)
	})

	api := r.Engine.Group("/api")
	{
		api.POST("/login", ah.Login)
		// router api for todo
		todoAPI := api.Group("/todo")
		todoAPI.Use(middleware.AuthUser(ur))
		{
			todoAPI.POST("/create", th.Create)

		}
	}
}
