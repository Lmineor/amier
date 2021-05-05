package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// ParsePageParams return pageNum and pageSize in the request params
func ParsePageParams(c *gin.Context) (pageNum, pageSize int) {
	var err error
	pageNum, err = strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err = strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 10
	}

	switch {
	case pageSize > 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}
	return
}

func ParsePoetId(c *gin.Context) (poetId string) {
	return c.Query("id")
}
