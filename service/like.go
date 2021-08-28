package service

import (
	"ziyue/global"
	"ziyue/model"
)

func GetPoems(params map[string]interface{}) (poems []model.Poem, total int64, err error) {
	limit := params["pageSize"].(int)
	offset := (params["pageNum"].(int) - 1) * limit
	desc := params["desc"].(bool)
	db := global.Z_DB.Model(&model.Poem{})
	err = db.Count(&total).Error
	if desc {
		err = db.Order("ilike desc").Limit(limit).Offset(offset).Find(&poems).Error
	} else {
		err = db.Limit(limit).Offset(offset).Find(&poems).Error
	}
	return
}
