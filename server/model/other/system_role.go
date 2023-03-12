// 自动生成模板SystemRole
package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// SystemRole 结构体
type SystemRole struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:角色名称;size:30;"`
      Code  string `json:"code" form:"code" gorm:"column:code;comment:角色代码;size:100;"`
      DataScope  *int `json:"dataScope" form:"dataScope" gorm:"column:data_scope;comment:数据范围（1：全部数据权限 2：自定义数据权限 3：本部门数据权限 4：本部门及以下数据权限 5：本人数据权限）;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态 (1正常 2停用);"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`
      CreatedBy  *int `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:19;"`
      UpdatedBy  *int `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:19;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
}


// TableName SystemRole 表名
func (SystemRole) TableName() string {
  return "system_role"
}

