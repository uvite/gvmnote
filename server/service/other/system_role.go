package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/other"
	otherReq "github.com/flipped-aurora/gin-vue-admin/server/model/other/request"
)

type SystemRoleService struct {
}

// CreateSystemRole 创建SystemRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemRoleService *SystemRoleService) CreateSystemRole(systemRole other.SystemRole) (err error) {
	err = global.GVA_DB.Create(&systemRole).Error
	return err
}

// DeleteSystemRole 删除SystemRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemRoleService *SystemRoleService) DeleteSystemRole(systemRole other.SystemRole) (err error) {
	err = global.GVA_DB.Delete(&systemRole).Error
	return err
}

// DeleteSystemRoleByIds 批量删除SystemRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemRoleService *SystemRoleService) DeleteSystemRoleByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]other.SystemRole{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSystemRole 更新SystemRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemRoleService *SystemRoleService) UpdateSystemRole(systemRole other.SystemRole) (err error) {
	err = global.GVA_DB.Save(&systemRole).Error
	return err
}

// GetSystemRole 根据id获取SystemRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemRoleService *SystemRoleService) GetSystemRole(id uint) (systemRole other.SystemRole, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&systemRole).Error
	return
}

// GetSystemRoleInfoList 分页获取SystemRole记录
// Author [piexlmax](https://github.com/piexlmax)
func (systemRoleService *SystemRoleService) GetSystemRoleInfoList(info otherReq.SystemRoleSearch) (list []other.SystemRole, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&other.SystemRole{})
	var systemRoles []other.SystemRole
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&systemRoles).Error
	return systemRoles, total, err
}
