package utils

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetPage(c *gin.Context, defPageNum, defPageLimit int) (int, int) {
	offset := 0
	p, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		p = defPageNum
	}
	l, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		l = defPageLimit
	}
	page := int(p)
	limit := int(l)
	if page > 0 {
		offset = (page - 1) * limit
	}
	return offset, limit
}
