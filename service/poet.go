package service

import (
	"errors"
	"ziyue/global"
	"ziyue/model"
	"ziyue/utils"

	"github.com/gookit/color"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPoetInfo(uuid string) (poet model.Poet, err error) {
	err = global.ZDB.Preload("Poems").Where("uuid = ?", uuid).First(&poet).Error
	return
}

func GetPoets(pageNum, pageSize int) (list []model.Poet, total int64, err error) {
	limit := pageSize
	offset := (pageNum - 1) * pageSize
	var poetList []model.Poet

	db := global.ZDB.Model(&model.Poet{})
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&poetList).Error
	return poetList, total, err
}

func CreatePoet(p model.Poet) (createdPoet model.Poet, err error) {
	var poet model.Poet
	if !errors.Is(global.ZDB.Where("poet = ? AND dynasty = ?", p.Poet, p.Dynasty).First(&poet).Error, gorm.ErrRecordNotFound) { // 判断诗人是否已存在
		return createdPoet, errors.New("该诗人已存在")
	}
	// 生成uuid 并存储
	p.UUID = utils.GeneratorUUID()
	err = global.ZDB.Omit(clause.Associations).Create(&p).Error
	return p, err
}

func GetPoetIdOrCreatePoet(poet, dynasty string) (uint, error) {
	var p model.Poet
	err := global.ZDB.Where("poet = ? AND dynasty = ?", poet, dynasty).First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		color.Errorf("ERROR: Cant not fetch poet %s's id, we create it.", poet)
		newPoet := &model.Poet{Poet: poet, Dynasty: dynasty}
		createdPoet, err := CreatePoet(*newPoet)
		return createdPoet.ID, err
	}
	return p.ID, err
}
func GetPoetUUID(pid uint) (uuid string, err error) {
	var p model.Poet
	err = global.ZDB.Where("id = ?", pid).First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		color.Infof("Poet %s not found", pid)
		return "", err
	}
	return p.UUID, nil

}

func UpdatePoet(p *model.Poet, uuid string) (poet *model.Poet, err error) {
	db := global.ZDB
	err = db.Where("uuid = ?", uuid).First(&poet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	poet.Poet = p.Poet
	poet.Dynasty = p.Dynasty
	poet.Descb = p.Descb
	db.Save(&poet)
	return poet, nil
}
