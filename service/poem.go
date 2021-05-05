package service

import (
	"fmt"
	"ziyue/global"
	"ziyue/model"
)

func GetPoetIntro(id string) (poet model.Poet, err error) {
	fmt.Println(id)
	err = global.ZDB.Where("id = ?", id).First(&poet).Error
	return
}

// func LunYuPaginate(c *gin.Context) {

// 	pageNum, pageSize := helper.ParsePageParams(c)

// 	result, err := new(model.Poem).List(pageNum, pageSize)

// 	if err != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"status": false,
// 			"data":   "",
// 			"msg":    err.Error(),
// 		})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{
// 		"status": true,
// 		"data":   result,
// 		"msg":    "查询成功",
// 	})

// }

// func CreateExaCustomer(e model.ExaCustomer) (err error) {
// 	err = global.GVA_DB.Create(&e).Error
// 	return err
// }

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: DeleteFileChunk
// //@description: 删除客户
// //@param: e model.ExaCustomer
// //@return: err error

// func DeleteExaCustomer(e model.ExaCustomer) (err error) {
// 	err = global.GVA_DB.Delete(&e).Error
// 	return err
// }

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: UpdateExaCustomer
// //@description: 更新客户
// //@param: e *model.ExaCustomer
// //@return: err error

// func UpdateExaCustomer(e *model.ExaCustomer) (err error) {
// 	err = global.GVA_DB.Save(e).Error
// 	return err
// }

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: GetExaCustomer
// //@description: 获取客户信息
// //@param: id uint
// //@return: err error, customer model.ExaCustomer

// func GetExaCustomer(id uint) (err error, customer model.ExaCustomer) {
// 	err = global.GVA_DB.Where("id = ?", id).First(&customer).Error
// 	return
// }

// //@author: [piexlmax](https://github.com/piexlmax)
// //@function: GetCustomerInfoList
// //@description: 分页获取客户列表
// //@param: sysUserAuthorityID string, info request.PageInfo
// //@return: err error, list interface{}, total int64

// func GetCustomerInfoList(sysUserAuthorityID string, info request.PageInfo) (err error, list interface{}, total int64) {
// 	limit := info.PageSize
// 	offset := info.PageSize * (info.Page - 1)
// 	db := global.GVA_DB.Model(&model.ExaCustomer{})
// 	var a model.SysAuthority
// 	a.AuthorityId = sysUserAuthorityID
// 	err, auth := GetAuthorityInfo(a)
// 	var dataId []string
// 	for _, v := range auth.DataAuthorityId {
// 		dataId = append(dataId, v.AuthorityId)
// 	}
// 	var CustomerList []model.ExaCustomer
// 	err = db.Where("sys_user_authority_id in ?", dataId).Count(&total).Error
// 	if err != nil {
// 		return err, CustomerList, total
// 	} else {
// 		err = db.Limit(limit).Offset(offset).Preload("SysUser").Where("sys_user_authority_id in ?", dataId).Find(&CustomerList).Error
// 	}
// 	return err, CustomerList, total
// }
