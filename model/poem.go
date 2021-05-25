package model

import (
	"ziyue/global"
)

type Poem struct {
	global.CommonModel
	Paragraphs string `json:"paragraphs" form:"paragraphs" gorm:"column:paragraphs;type:text;comment:诗歌内容"`
	Poem       string `json:"poem" form:"poem" gorm:"column:poem;comment:poem title"`
	Like       uint   `json:"ilike" gorm:"column:ilike;comment:ilike;default:0"`
	PoetID     uint   `json:"poet_id" ` // 外键
}

func (Poem) TableName() string {
	return "poems"
}
