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
	UUID      string         `json:"uuid" gorm:"column:uuid;comment:uuid"`
}
