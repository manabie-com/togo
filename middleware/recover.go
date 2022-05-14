package middleware

import (
	"github.com/japananh/togo/common"
	"github.com/japananh/togo/component"

	"github.com/gin-gonic/gin"
)

// Recover needs AppContext due to 2 reasons below
// - Log error to DB
// - Get enviroment settings (prod/staging/dev)
func Recover(appCtx component.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				// if error is an AppError
				if appErr, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appErr.StatusCode, appErr)
					// Gin has its own `Recover`, that wraps our `Recover`
					// Gin can dumb your error to the terminal when we call `panic` here.
					// It makes dev easier to trace bugs.
					// Call `return` here won't dumb error in the terminal
					panic(err)
				}

				appErr := common.ErrInternal((err.(error)))
				c.AbortWithStatusJSON(appErr.StatusCode, appErr)
				panic(err)
			}
		}()

		c.Next()
	}
}
