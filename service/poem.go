package service

import (
	"errors"
	"strings"
	"ziyue/global"
	"ziyue/model"
	"ziyue/model/request"
	"ziyue/utils"

	"gorm.io/gorm"
)

func CreatePoem(p *model.Poem, poet, dynasty string) (err error) {
	poetId, _ := GetPoetIdOrCreatePoet(poet, dynasty)
	p.PoetID = poetId
	p.UUID = utils.GeneratorUUID()
	err = global.ZDB.Create(&p).Error
	return
}

func UpdatePoem(p *request.Poem, uuid string) (poem *model.Poem, err error) {
	db := global.ZDB
	err = db.Where("uuid = ?", uuid).First(&poem).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	poem.Poem = p.Poem
	poem.Paragraphs = strings.Join(p.Paragraphs, "|")
	db.Save(&poem)
	return poem, nil
}

func GetPoem(uuid string) (poem model.Poem, err error) {
	err = global.ZDB.Where("uuid = ?", uuid).First(&poem).Error
	return
}
