package model

import (
	"ziyue/global"
)

type Poet struct {
	global.CommonModel
	Poet    string `json:"poet" form:"poet" gorm:"comment:诗人人名"`
	Dynasty string `json:"dynasty" form:"dynasty" gorm:"comment:朝代"`
	Descb   string `json:"descb" form:"descb" gorm:"type:text;comment:诗人简介"`
	Poems   []Poem `json:"poems"`
}

type Poem struct {
	global.CommonModel
	Paragraphs string `json:"paragraphs" form:"paragraphs" gorm:"column:paragraphs;type:text;comment:诗歌内容"`
	Poem       string `json:"poem" form:"poem" gorm:"column:poem;comment:poem title"`
	Like       uint   `json:"ilike" gorm:"column:ilike;comment:ilike;default:0"`
	PoetID     uint   `json:"poet_id" ` // 外键
}

func (Poet) TableName() string {
	return "poets"
}

func (Poem) TableName() string {
	return "poems"
}
