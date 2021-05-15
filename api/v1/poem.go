package api

import (
	"net/http"
	"ziyue/model"
	"ziyue/model/response"
	"ziyue/service"
	"ziyue/utils"

	"github.com/gin-gonic/gin"
	"github.com/gookit/color"
)

type createPoemStruct struct {
	Poem       string `json:"poem"`
	Paragraphs string `json:"paragraphs"`
	Poet       string `json:"poet"`
	Dynasty    string `json:"dynasty"`
}

func GetPoet(c *gin.Context) {
	uuid := utils.ParsePoetUUId(c)
	if uuid == "" {
		if poetList, total, err := service.GetPoets(c); err != nil {
			response.FailWithMessage("error", c)
		} else {
			response.OkWithData(response.PoetsResponse{Poets: poetList, Total: total}, c)
		}
	} else {
		if poet, err := service.GetPoetInfo(uuid); err != nil {
			response.FailWithMessage("没有这个诗人的记录", c)
		} else {
			response.OkWithDetailed(response.PoetResponse{Poet: poet}, "success", c)
		}
	}

}

func GetPoemLike(c *gin.Context) {
	pageNum, pageSize := utils.ParsePageParams(c)
	like, err := service.GetLikePoemList(pageNum, pageSize)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "error",
			"msg":    err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   like,
			"msg":    "success",
		})
	}
}

// CreatePoet according to request
func CreatePoet(c *gin.Context) {
	var poet model.Poet
	_ = c.ShouldBindJSON(&poet)
	// if err := utils.Verify(R, utils.RegisterVerify); err != nil {
	// 	response.FailWithMessage(err.Error(), c)
	// 	return
	// } else {

	// }
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
	newPoem := &model.Poem{Paragraphs: poem.Paragraphs, Poem: poem.Poem}
	err := service.CreatePoem(newPoem, poem.Poet, poem.Dynasty)
	if err != nil {
		color.Errorf("Create poem failed for %s", err)
	}
	response.Ok(c)
}
