package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/other"
	otherReq "github.com/flipped-aurora/gin-vue-admin/server/model/other/request"
)

type SettingGenerateTablesService struct {
}

// CreateSettingGenerateTables 创建SettingGenerateTables记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateTablesService *SettingGenerateTablesService) CreateSettingGenerateTables(settingGenerateTables other.SettingGenerateTables) (err error) {
	err = global.GVA_DB.Create(&settingGenerateTables).Error
	return err
}

// DeleteSettingGenerateTables 删除SettingGenerateTables记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateTablesService *SettingGenerateTablesService)DeleteSettingGenerateTables(settingGenerateTables other.SettingGenerateTables) (err error) {
	err = global.GVA_DB.Delete(&settingGenerateTables).Error
	return err
}

// DeleteSettingGenerateTablesByIds 批量删除SettingGenerateTables记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateTablesService *SettingGenerateTablesService)DeleteSettingGenerateTablesByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]other.SettingGenerateTables{},"id in ?",ids.Ids).Error
	return err
}

// UpdateSettingGenerateTables 更新SettingGenerateTables记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateTablesService *SettingGenerateTablesService)UpdateSettingGenerateTables(settingGenerateTables other.SettingGenerateTables) (err error) {
	err = global.GVA_DB.Save(&settingGenerateTables).Error
	return err
}

// GetSettingGenerateTables 根据id获取SettingGenerateTables记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateTablesService *SettingGenerateTablesService)GetSettingGenerateTables(id uint) (settingGenerateTables other.SettingGenerateTables, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&settingGenerateTables).Error
	return
}

// GetSettingGenerateTablesInfoList 分页获取SettingGenerateTables记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateTablesService *SettingGenerateTablesService)GetSettingGenerateTablesInfoList(info otherReq.SettingGenerateTablesSearch) (list []other.SettingGenerateTables, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&other.SettingGenerateTables{})
    var settingGenerateTabless []other.SettingGenerateTables
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.StartCreatedAt !=nil && info.EndCreatedAt !=nil {
     db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }

	err = db.Limit(limit).Offset(offset).Find(&settingGenerateTabless).Error
	return  settingGenerateTabless, total, err
}
