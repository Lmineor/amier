package model

import "ziyue/global"

type SongCi struct {
	global.CommonModel
	Paragraphs string `json:"paragraphs" form:"paragraphs" gorm:"comment:宋词内容"`
	Rhythmic   string `json:"rhythmic" form:"rhythmic" gorm:"comment:词牌名"`
	Poetid     uint   `json:"poetid" form:"poetid" gorm:"foreignKey:CiPoet;  comment:词作者"`
	Uid        string `json:"u_id" form:"u_id" gorm:"comment:uid"`
}

type CiPoet struct {
	global.CommonModel
	Longdesc  string `json:"Longdesc" gorm:"comment: 详细介绍"`
	Shortdesc string `json:"shortdesc" gorm:"comment:简介"`
	Poet      string `json:"poet" gorm:"comment:词作者"`
	Ci        string `json:"ci" gorm:"comment:ci"`
}
