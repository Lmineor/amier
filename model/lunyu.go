package model

import "ziyue/global"

type Lunyu struct {
	global.CommonModel
	Paragraphs string `json:"paragraphs" form:"paragraphs" gorm:"comment:论语内容"`
	Chapter    string `json:"chapter" form:"chapter" gorm:"comment:论语标题"`
	Uid        string `json:"u_id" form:"u_id" gorm:"comment:uid"`
}
