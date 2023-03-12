package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/other"
	otherReq "github.com/flipped-aurora/gin-vue-admin/server/model/other/request"
)

type SettingGenerateColumnsService struct {
}

// CreateSettingGenerateColumns 创建SettingGenerateColumns记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateColumnsService *SettingGenerateColumnsService) CreateSettingGenerateColumns(settingGenerateColumns other.SettingGenerateColumns) (err error) {
	err = global.GVA_DB.Create(&settingGenerateColumns).Error
	return err
}

// DeleteSettingGenerateColumns 删除SettingGenerateColumns记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateColumnsService *SettingGenerateColumnsService) DeleteSettingGenerateColumns(settingGenerateColumns other.SettingGenerateColumns) (err error) {
	err = global.GVA_DB.Delete(&settingGenerateColumns).Error
	return err
}

// DeleteSettingGenerateColumnsByIds 批量删除SettingGenerateColumns记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateColumnsService *SettingGenerateColumnsService) DeleteSettingGenerateColumnsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]other.SettingGenerateColumns{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSettingGenerateColumns 更新SettingGenerateColumns记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateColumnsService *SettingGenerateColumnsService) UpdateSettingGenerateColumns(settingGenerateColumns other.SettingGenerateColumns) (err error) {
	err = global.GVA_DB.Save(&settingGenerateColumns).Error
	return err
}

// GetSettingGenerateColumns 根据id获取SettingGenerateColumns记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateColumnsService *SettingGenerateColumnsService) GetSettingGenerateColumns(id uint) (settingGenerateColumns other.SettingGenerateColumns, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&settingGenerateColumns).Error
	return
}

// GetSettingGenerateColumnsInfoList 分页获取SettingGenerateColumns记录
// Author [piexlmax](https://github.com/piexlmax)
func (settingGenerateColumnsService *SettingGenerateColumnsService) GetSettingGenerateColumnsInfoList(info otherReq.SettingGenerateColumnsSearch) (list []other.SettingGenerateColumns, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&other.SettingGenerateColumns{})
	var settingGenerateColumnss []other.SettingGenerateColumns
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&settingGenerateColumnss).Error
	return settingGenerateColumnss, total, err
}
