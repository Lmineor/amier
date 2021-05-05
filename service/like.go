package service

import (
	"ziyue/global"
	"ziyue/model"
)

func GetLikePoemList(pageNum, pageSize int) (like []model.LikePoem, err error) {
	// pageNum, pageSize := helper.ParsePageParams(c)
	db := global.ZDB
	limit := pageSize //TODO verify the pagesize
	offset := pageSize * (pageNum - 1)
	// var like []model.LikePoem
	err = db.Order("i_like DESC").Limit(limit).Offset(offset).Preload("Poem").Find(&like).Error
	return
}

// func (t *TangShi) List(pageNum int, pageSize int, keyword string) (*PageResult, error) {

// 	var r []TangShi
// 	var maps, or map[string]interface{}

// 	if keyword != "" {
// 		maps = map[string]interface{}{"title": keyword}
// 		or = map[string]interface{}{"author": keyword}
// 	}

// 	result, err := getPaginateData(t, pageNum, pageSize, maps, or, func(q *gorm.DB) error { return q.Find(&r).Error })

// 	if err != nil {
// 		return result, err
// 	}
// 	result.List = r

// 	return result, nil
// }
