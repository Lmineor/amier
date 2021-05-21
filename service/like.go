package service

import (
	"ziyue/global"
	"ziyue/model"
)

func GetLikePoemList(pageNum, pageSize int) (data interface{}, err error) {
	limit := pageSize
	offset := (pageNum - 1) * pageSize
	var poemLikeList []model.LikePoem

	Like := make(map[string]interface{})
	var total int64

	db := global.ZDB.Preload("Poem").Model(&model.LikePoem{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&poemLikeList).Error
	Like["total"] = total
	Like["like"] = poemLikeList
	return Like, err
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
