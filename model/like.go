package model

import "ziyue/global"

type LikePoem struct {
	global.Like
	Poem Poem `json:"poem" form:"peom" gorm:"ForeignKey:Uid;AssociationForeignKey:Uid;references:Uid;comment:诗经内容"`
}

func (LikePoem) TableName() string {
	return "like_poem"
}
