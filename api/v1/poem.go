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

// func GetPoemLike(c *gin.Context) {
// 	pageNum, pageSize := utils.ParsePageParams(c)
// 	like, err := service.GetLikePoemList(pageNum, pageSize)
// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": false,
// 			"data":   "error",
// 			"msg":    err.Error(),
// 		})
// 	} else {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": true,
// 			"data":   like,
// 			"msg":    "success",
// 		})
// 	}
// }

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
		service.GetPoems()
		response.Ok(c)
	} else {
		poem, err := service.GetPoem(uuid)
		if err != nil {
			response.FailWithMessage("无记录", c)
		} else {
			pUUID, _ := service.GetPoetUUID(poem.ID)
			response.OkWithData(response.PoemResponse{
				Poem:       poem.Poem,
				UUID:       uuid,
				Paragraphs: strings.Split(poem.Paragraphs, "|"),
				PoetUUID:   pUUID,
			}, c)
		}
	}
}

func GetLike(c *gin.Context) {
	pageNum, pageSize := utils.ParsePageParams(c)
	mode := utils.GetLikeMode(c)
	switch mode {
	case "poem":
		data, _ := service.GetLikePoemList(pageNum, pageSize)
		response.OkWithData(data, c)
	default:
		data, _ := service.GetLikePoemList(pageNum, pageSize)
		response.OkWithData(data, c)
	}

}
