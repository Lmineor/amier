package api

import (
	"ziyue/model"
	"ziyue/model/response"
	"ziyue/service"
	"ziyue/utils"

	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

// CreatePoet from request
func CreatePoet(c *gin.Context) {
	var poet model.Poet
	_ = c.ShouldBindJSON(&poet)
	p := &model.Poet{Poet: poet.Poet, Dynasty: poet.Dynasty, Descb: poet.Descb}
	_, err := service.CreatePoet(*p)
	if err != nil {
		color.Error.Renderln(err)
		response.FailWithMessage(err.Error(), c)
	} else {
		response.Ok(c)
	}

}

func GetPoet(c *gin.Context) {
	uuid := utils.ParseReqUUId(c)
	if uuid == "" {
		pageNum, pageSize := utils.ParsePageParams(c)
		if poetList, total, err := service.GetPoets(pageNum, pageSize); err != nil {
			response.FailWithMessage("error", c)
		} else {
			response.OkWithData(response.PoetsResponse{Poets: poetList, Total: total}, c)
		}
	} else {
		if poet, err := service.GetPoetInfo(uuid); err != nil {
			color.Info.Print(err)
			response.FailWithMessage("没有这个诗人的记录", c)
		} else {
			poems := make([]string, 0)
			for _, poem := range poet.Poems {
				poems = append(poems, poem.UUID)
			}
			response.OkWithData(response.PoetResponse{
				Poet:    poet.Poet,
				Dynasty: poet.Dynasty,
				Descb:   poet.Descb,
				UUID:    poet.UUID,
				Poems:   poems,
			}, c)
		}
	}

}

func UpdatePoet(c *gin.Context) {
	var poet model.Poet
	c.ShouldBindJSON(&poet)
	uuid := utils.ParseReqUUId(c)
	_, err := service.UpdatePoet(&poet, uuid)
	color.Debug.Printf("updated peom's uuid is: %s\n", uuid)
	if err != nil {
		response.FailWithMessage("记录不存在", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}
