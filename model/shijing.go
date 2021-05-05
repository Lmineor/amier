package model

import "ziyue/global"

type Shijing struct {
	global.CommonModel
	Poem    string `json:"poem" form:"poem" gorm:"comment:诗经内容"`
	Chapter string `json:"chapter" form:"chapter" gorm:"comment:章"`
	Section string `json:"section" form:"section" gorm:"comment:篇"`
	Content string `json:"content" form:"content" gorm:"comment:诗经内容"`
	Uid     string `json:"u_id" form:"u_id" gorm:"comment:uid"`
}
