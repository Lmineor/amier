package api

import (
	"strings"
	"ziyue/model"
	"ziyue/model/response"
	"ziyue/service"
	"ziyue/utils"

	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

type createPoemStruct struct {
	Poem       string   `json:"poem"`
	Paragraphs []string `json:"paragraphs"`
	Poet       string   `json:"poet"`
	Dynasty    string   `json:"dynasty"`
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

func CreatePoem(c *gin.Context) {
	var poem createPoemStruct
	c.ShouldBindJSON(&poem)
	newPoem := &model.Poem{Paragraphs: strings.Join(poem.Paragraphs, "|"), Poem: poem.Poem}
	err := service.CreatePoem(newPoem, poem.Poet, poem.Dynasty)
	if err != nil {
		color.Errorf("Create poem failed for %s", err)
	}
	response.Ok(c)
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

func GetLikes(c *gin.Context) {
	var likes []model.Poem
	var total int64
	likePoems := make([]response.PoemResponse, 0)
	respMap := make(map[string]interface{})

	pageNum, pageSize := utils.ParsePageParams(c)
	mode := utils.GetLikeMode(c)

	switch mode {
	case "poem":
		likes, total, _ = service.GetLikePoems(pageNum, pageSize)
		respMap["total"] = total
		for _, poem := range likes {
			pUUID, _ := service.GetPoetUUID(poem.ID)
			likePoems = append(likePoems, *utils.ParsePoemSplit(&poem, pUUID))
		}
		respMap["pems"] = likePoems

	default:
		likes, total, _ = service.GetLikePoems(pageNum, pageSize)
		for _, poem := range likes {
			pUUID, _ := service.GetPoetUUID(poem.ID)
			likePoems = append(likePoems, *utils.ParsePoemSplit(&poem, pUUID))
		}
	}

	respMap["pems"] = likePoems
	respMap["total"] = total
	response.OkWithData(respMap, c)
}

func GetPoems(c *gin.Context) {
	GetLikes(c)
}
