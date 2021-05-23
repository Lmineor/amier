package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
	"ziyue/model"
	"ziyue/model/response"

	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

const (
	PomeLike    = 1 // 获取poem的like
	ShijingLike = 2 //获取诗经的like
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

func ParsePoemSplit(p *model.Poem, poetUUID string) *response.PoemResponse {
	return &response.PoemResponse{
		Poem:       p.Poem,
		UUID:       p.UUID,
		Paragraphs: strings.Split(p.Paragraphs, "|"),
		PoetUUID:   poetUUID,
		Like:       p.Like,
	}
}
