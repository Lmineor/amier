package global

import "gorm.io/gorm"

type CommonModel struct {
	gorm.Model
	Uuid string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:uuid"`
}

type Like struct {
	CommonModel
	Ilike uint   `json:"ilike" gorm:"column:i_like;comment:like"`
	Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:uuid"`
}
