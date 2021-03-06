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
)

func CreatePoem(c *gin.Context) {
	var poem request.Poem
	c.ShouldBindJSON(&poem)
	newPoem := &model.Poem{Paragraphs: strings.Join(poem.Paragraphs, "|"), Poem: poem.Poem}
	err := service.CreatePoem(newPoem, poem.Poet, poem.Dynasty)
	if err != nil {
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
			response.FailWithMessage("No record!", c)
		} else {
			pUUID, _ := service.GetPoetUUID(poem.ID)
			response.OkWithData(
				&response.PoemResponse{
					Poem:       poem.Poem,
					UUID:       poem.UUID,
					Paragraphs: strings.Split(poem.Paragraphs, "|"),
					PoetUUID:   pUUID,
					Like:       poem.Like},
				c)
		}
	}
}

func UpdatePoem(c *gin.Context) {
	var poem request.Poem
	c.ShouldBindJSON(&poem)
	uuid := utils.ParseReqUUId(c)
	_, err := service.UpdatePoem(&poem, uuid)
	global.Z_LOG.Info("update poem", zap.String("poem", uuid))
	if err != nil {
		response.FailWithMessage("记录不存在", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

}

// GetPoems return poems info
func GetPoems(c *gin.Context) {
	var poems []model.Poem
	var total int64
	parsedPoems := make([]response.PoemResponse, 0) //声明并完成切片的初始化
	respMap := make(map[string]interface{})         //声明并完成map的初始化

	params, _ := utils.ParseParams(c)
	poems, total, _ = service.GetPoems(params)
	respMap["total"] = total
	for _, poem := range poems {
		pUUID, _ := service.GetPoetUUID(poem.ID)
		parsedPoems = append(parsedPoems, *utils.ParsePoemSplit(&poem, pUUID))
	}
	respMap["poems"] = parsedPoems
	respMap["total"] = total
	response.OkWithData(respMap, c)
}
