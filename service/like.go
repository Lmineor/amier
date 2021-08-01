package service

import (
	"fmt"
	"ziyue/global"
	"ziyue/model"
)

func GetLikePoems(pageNum, pageSize int) (poemsLike []model.Poem, total int64, err error) {
	limit := pageSize
	offset := (pageNum - 1) * pageSize
	fmt.Printf("limit is %d, offset is %d", limit, offset)
	db := global.Z_DB.Model(&model.Poem{})
	err = db.Count(&total).Error
	err = db.Order("ilike desc").Limit(limit).Offset(offset).Find(&poemsLike).Error
	return
}
