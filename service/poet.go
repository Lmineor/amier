package service

import (
	"errors"
	"go.uber.org/zap"
	"ziyue/global"
	"ziyue/model"
	"ziyue/utils"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func GetPoetInfo(uuid string) (poet model.Poet, err error) {
	err = global.Z_DB.Preload("Poems").Where("uuid = ?", uuid).First(&poet).Error
	return
}

func GetPoets(params map[string]interface{}) (list []model.Poet, total int64, err error) {
	limit := params["pageSize"].(int)
	offset := (params["pageNum"].(int) - 1) * limit
	showPoems := params["showPoems"].(bool)
	global.Z_LOG.Info("msg:", zap.Any("showpoems", showPoems))
	db := global.Z_DB.Model(&model.Poet{})
	err = db.Count(&total).Error
	if showPoems {
		err = db.Limit(limit).Offset(offset).Preload("Poems").Find(&list).Error
	} else {
		err = db.Limit(limit).Offset(offset).Find(&list).Error
	}

	return
}

func CreatePoet(p model.Poet) (createdPoet model.Poet, err error) {
	if !errors.Is(global.Z_DB.Where("poet = ? AND dynasty = ?", p.Poet, p.Dynasty).First(&createdPoet).Error, gorm.ErrRecordNotFound) { // 判断诗人是否已存在
		return createdPoet, errors.New("该诗人已存在")
	}
	// 生成uuid 并存储
	p.UUID = utils.GeneratorUUID()
	err = global.Z_DB.Omit(clause.Associations).Create(&p).Error
	return p, err
}

func GetPoetIdOrCreatePoet(poet, dynasty string) (uint, error) {
	var p model.Poet
	err := global.Z_DB.Where("poet = ? AND dynasty = ?", poet, dynasty).First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newPoet := &model.Poet{Poet: poet, Dynasty: dynasty}
		createdPoet, err := CreatePoet(*newPoet)
		return createdPoet.ID, err
	}
	return p.ID, err
}
func GetPoetUUID(pid uint) (uuid string, err error) {
	var p model.Poet
	err = global.Z_DB.Where("id = ?", pid).First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		global.Z_LOG.Debug("")
		return "", err
	}
	return p.UUID, nil

}

func UpdatePoet(p *model.Poet, uuid string) (poet *model.Poet, err error) {
	db := global.Z_DB
	err = db.Where("uuid = ?", uuid).First(&poet).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}
	poet.Poet = p.Poet
	poet.Dynasty = p.Dynasty
	poet.Describe = p.Describe
	db.Save(&poet)
	return poet, nil
}
