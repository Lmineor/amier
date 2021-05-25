package model

import "ziyue/global"

type Poet struct {
	global.CommonModel
	Poet    string `json:"poet" form:"poet" gorm:"comment:诗人人名"`
	Dynasty string `json:"dynasty" form:"dynasty" gorm:"comment:朝代"`
	Descb   string `json:"descb" form:"descb" gorm:"type:text;comment:诗人简介"`
	Poems   []Poem `json:"poems"`
}

func (Poet) TableName() string {
	return "poets"
}
