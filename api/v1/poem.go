package api

import (
	"go.uber.org/zap"
	"strings"
	"ziyue/global"
	"ziyue/model"
	"ziyue/model/request"
	"ziyue/model/response"
	"ziyue/service"
	"ziyue/utils"

	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

func CreatePoem(c *gin.Context) {
	var poem request.Poem
	c.ShouldBindJSON(&poem)
	newPoem := &model.Poem{Paragraphs: strings.Join(poem.Paragraphs, "|"), Poem: poem.Poem}
	err := service.CreatePoem(newPoem, poem.Poet, poem.Dynasty)
	if err != nil {
		color.Errorf("Create poem failed for %s", err)
		global.Z_LOG.Error("创建poem失败！", zap.Any("err", err))
		response.Fail(c)
	} else {
		response.Ok(c)
	}
}

func GetPoem(c *gin.Context) {
	uuid := utils.ParseReqUUId(c)

	if uuid == "" {
		GetPoems(c)
	} else {
		poem, err := service.GetPoem(uuid)
		if err != nil {
			response.FailWithMessage("无记录", c)
		} else {
			pUUID, _ := service.GetPoetUUID(poem.ID)
			response.OkWithData(utils.ParsePoemSplit(&poem, pUUID), c)
		}
	}
}

func UpdatePoem(c *gin.Context) {
	var poem request.Poem
	c.ShouldBindJSON(&poem)
	uuid := utils.ParseReqUUId(c)
	_, err := service.UpdatePoem(&poem, uuid)
	color.Debug.Printf("updated peom's uuid is: %s\n", uuid)
	if err != nil {
		response.FailWithMessage("记录不存在", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

}

func GetPoems(c *gin.Context) {
	GetLikes(c)
}
