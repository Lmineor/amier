package global

import (
	"time"

	"gorm.io/gorm"
)

type CommonModel struct {
	ID        uint           `json:"-" gorm:"primarykey"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `json:"-"`
	UUID      string         `json:"uuid" form:"uuid" gorm:"column:uuid;comment:uuid"`
}

type Like struct {
	Ilike uint `json:"ilike" gorm:"column:like;comment:like"`
}
