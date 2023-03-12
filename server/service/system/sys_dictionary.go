package system

import (
	"errors"
	request2 "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"gorm.io/gorm"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionary
//@description: 创建字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

type DictionaryService struct{}

func (dictionaryService *DictionaryService) CreateSysDictionary(sysDictionary system.SysDictionary) (err error) {
	if (!errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "code = ?", sysDictionary.Code).Error, gorm.ErrRecordNotFound)) {
		return errors.New("存在相同的type，不允许创建")
	}
	err = global.GVA_DB.Create(&sysDictionary).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: DeleteSysDictionary
//@description: 删除字典数据
//@param: sysDictionary model.SysDictionary
//@return: err error

func (dictionaryService *DictionaryService) DeleteSysDictionary(sysDictionary system.SysDictionary) (err error) {
	err = global.GVA_DB.Where("id = ?", sysDictionary.ID).Preload("SysDictionaryDetails").First(&sysDictionary).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return errors.New("请不要搞事")
	}
	if err != nil {
		return err
	}
	err = global.GVA_DB.Delete(&sysDictionary).Error
	if err != nil {
		return err
	}

	if sysDictionary.SysDictionaryDetails != nil {
		return global.GVA_DB.Where("sys_dictionary_id=?", sysDictionary.ID).Delete(sysDictionary.SysDictionaryDetails).Error
	}
	return
}

func (dictionaryService *DictionaryService) DeleteSysDictionarys(ids request2.IdsReq) (err error) {

	for _, id := range ids.Ids {

		var sysDictionary system.SysDictionary
		err = global.GVA_DB.Where("id = ?", id).Preload("SysDictionaryDetails").First(&sysDictionary).Error
		if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("请不要搞事")
		}
		if err != nil {
			return err
		}
		err = global.GVA_DB.Delete(&sysDictionary).Error
		if err != nil {
			return err
		}

		if sysDictionary.SysDictionaryDetails != nil {
			return global.GVA_DB.Where("sys_dictionary_id=?", sysDictionary.ID).Delete(sysDictionary.SysDictionaryDetails).Error
		}
	}

	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: UpdateSysDictionary
//@description: 更新字典数据
//@param: sysDictionary *model.SysDictionary
//@return: err error

func (dictionaryService *DictionaryService) UpdateSysDictionary(sysDictionary *system.SysDictionary) (err error) {
	var dict system.SysDictionary
	sysDictionaryMap := map[string]interface{}{
		"Name":   sysDictionary.Name,
		"Code":   sysDictionary.Code,
		"Status": sysDictionary.Status,
		"Desc":   sysDictionary.Desc,
	}
	db := global.GVA_DB.Where("id = ?", sysDictionary.ID).First(&dict)
	if dict.Code != sysDictionary.Code {
		if !errors.Is(global.GVA_DB.First(&system.SysDictionary{}, "code = ?", sysDictionary.Code).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同的type，不允许创建")
		}
	}
	err = db.Updates(sysDictionaryMap).Error
	return err
}

//func (dictionaryService *DictionaryService) ChangeSort(req *system.SysDictionary) (err error) {
//
//	return global.GVA_DB.Model(&system.SysDictionary{}).
//		Select("updated_at", "sort").
//		Where("id=?", req.ID).
//		Updates(map[string]interface{}{
//			"updated_at": time.Now(),
//			"sort":       req.Sort,
//		}).Error
//}

func (dictionaryService *DictionaryService) ChangeStatus(req system.SysDictionary) (err error) {
	return global.GVA_DB.Model(&system.SysDictionary{}).
		Select("updated_at", "status").
		Where("id=?", req.ID).
		Updates(map[string]interface{}{
			"updated_at": time.Now(),
			"status":     req.Status,
		}).Error
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetSysDictionary
//@description: 根据id或者type获取字典单条数据
//@param: Type string, Id uint
//@return: err error, sysDictionary model.SysDictionary

func (dictionaryService *DictionaryService) GetSysDictionary(Code string, Id uint, status int) (sysDictionary system.SysDictionary, err error) {

	err = global.GVA_DB.Where("(code = ? OR id = ?) ", Code, Id).Preload("SysDictionaryDetails", "status = ?", true).First(&sysDictionary).Error
	return
}

// 获取多条
func (dictionaryService *DictionaryService) GetSysDictionarys(Codes request2.CodesReq) (sysDictionary system.SysDictionary, err error) {
	var flag = true

	err = global.GVA_DB.Where("(codes   in (?)) and status = ?", Codes, flag).Preload("SysDictionaryDetails", "status = ?", true).First(&sysDictionary).Error
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@function: GetSysDictionaryInfoList
//@description: 分页获取字典列表
//@param: info request.SysDictionarySearch
//@return: err error, list interface{}, total int64

func (dictionaryService *DictionaryService) GetSysDictionaryInfoList(info request.SysDictionarySearch) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&system.SysDictionary{})
	var sysDictionarys []system.SysDictionary
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.Name != "" {
		db = db.Where("`name` LIKE ?", "%"+info.Name+"%")
	}
	if info.Code != "" {
		db = db.Where("`code` LIKE ?", "%"+info.Code+"%")
	}
	if info.Status != 0 {
		db = db.Where("`status` = ?", info.Status)
	}
	if info.Desc != "" {
		db = db.Where("`desc` LIKE ?", "%"+info.Desc+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&sysDictionarys).Error
	return sysDictionarys, total, err
}
