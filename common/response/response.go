package response

import "github.com/gin-gonic/gin"

func Sucess(data interface{}) gin.H {
	return gin.H{"data": data}
}

func Failure(message interface{}) gin.H {
	return gin.H{"message": message}
}
