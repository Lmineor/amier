package api

import (
	"net/http"
	"ziyue/service"
	"ziyue/utils"

	"github.com/gin-gonic/gin"
)

func GetPoems(c *gin.Context) {
	// var a string

	poetID := utils.ParsePoetId(c)
	if poet, err := service.GetPoetIntro(poetID); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": true,
			"data":   poet,
			"msg":    "success",
		})
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
