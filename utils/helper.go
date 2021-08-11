package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	"io/ioutil"
	"strconv"
)

const (
	PomeLike    = 1 // 获取poem的like
	ShijingLike = 2 //获取诗经的like
)

// ParseParams return pageNum, pageSize and showPoems in the request params
func ParseParams(c *gin.Context) (pageNum, pageSize int, showPoems bool) {
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

	showPoems_ := c.Query("showPoems")
	switch showPoems_ {
	case "true", "True", "1", "yes":
		showPoems = true
	default:
		showPoems = false
	}
	return
}

func ParseReqUUId(c *gin.Context) (uuid string) {
	return c.Query("uuid")
}

func GetReqBody(c *gin.Context) string {
	body, _ := ioutil.ReadAll(c.Request.Body)
	color.Info.Println(string(body))
	color.Info.Printf("format is %T\n", string(body))
	return "123"
}

func GetLikeMode(c *gin.Context) string {
	mode := c.Query("type")
	switch mode {
	case "shijing":
		return mode
	case "poem":
		return mode
	case "lunyu":
		return mode
	default:
		return "poem"
	}
}
