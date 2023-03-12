// 自动生成模板SysDictionary
package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 如果含有time.Time 请自行import time包
type SysDictionary struct {
	global.GVA_MODEL
	Name                 string                `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`   // 字典名（中）
	Code                 string                `json:"code" form:"code" gorm:"column:code;comment:字典名（英）"`   // 字典名（英）
	Status               int                 `json:"status" form:"status" gorm:"column:status;comment:状态"` // 状态
	Desc                 string                `json:"remark" form:"desc" gorm:"column:desc;comment:描述"`     // 描述
	SysDictionaryDetails []SysDictionaryDetail `json:"sysDictionaryDetails" form:"sysDictionaryDetails"`
}

func (SysDictionary) TableName() string {
	return "system_dict_type"
}
