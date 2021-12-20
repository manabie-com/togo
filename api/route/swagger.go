package route

import (
	"github.com/gin-gonic/gin"
	_ "github.com/manabie-com/togo/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func RegisterSwagger(r *gin.RouterGroup) {
	r.GET(swagger, ginSwagger.WrapHandler(swaggerFiles.Handler))
}
