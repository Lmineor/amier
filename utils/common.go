package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
	uuid "github.com/satori/go.uuid"
	"io/ioutil"
	"strconv"
)

const (
	PomeLike    = 1 // 获取poem的like
	ShijingLike = 2 //获取诗经的like
)

// ParseParams return pageNum, pageSize and showPoems in the request params
func ParseParams(c *gin.Context) (params map[string]interface{}, err error) {
	params = make(map[string]interface{})
	var pageNum, pageSize int
	var showPoems, desc bool
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
	params["pageSize"] = pageSize
	params["pageNum"] = pageNum

	showPoems = trans2Bool(c.Query("showPoems"))
	params["showPoems"] = showPoems
	desc = trans2Bool(c.Query("desc"))
	params["desc"] = desc
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

// getBool judge bs  boolstring whether is bool
func trans2Bool(bs string) bool {
	switch bs {
	case "true", "True", "1", "yes":
		return true
	default:
		return false
	}
}

func GeneratorUUID() string {
	id := uuid.NewV4()
	return id.String()
}
