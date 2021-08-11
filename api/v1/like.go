package api

//func GetLikes(c *gin.Context) {
//	var likes []model.Poem
//	var total int64
//	likePoems := make([]response.PoemResponse, 0)
//	respMap := make(map[string]interface{})
//
//	pageNum, pageSize, _ := utils.ParseParams(c)
//	mode := utils.GetLikeMode(c)
//
//	switch mode {
//	case "poem":
//		likes, total, _ = service.GetLikePoems(pageNum, pageSize)
//		respMap["total"] = total
//		for _, poem := range likes {
//			pUUID, _ := service.GetPoetUUID(poem.ID)
//			likePoems = append(likePoems, *utils.ParsePoemSplit(&poem, pUUID))
//		}
//		respMap["poems"] = likePoems
//
//	default:
//		likes, total, _ = service.GetLikePoems(pageNum, pageSize)
//		for _, poem := range likes {
//			pUUID, _ := service.GetPoetUUID(poem.ID)
//			likePoems = append(likePoems, *utils.ParsePoemSplit(&poem, pUUID))
//		}
//	}
//
//	respMap["pems"] = likePoems
//	respMap["total"] = total
//	response.OkWithData(respMap, c)
//}
