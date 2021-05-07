package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

func TokenUserID(context *gin.Context) string {
	val, _ := context.Get("user_id")
	return fmt.Sprintf("%v", val)
}

func TokenMaxTodo(context *gin.Context) int {
	val, _ := context.Get("max_todo")
	maxTodo, _ :=  strconv.Atoi(fmt.Sprintf("%v", val))
	return maxTodo
}
