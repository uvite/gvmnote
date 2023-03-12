// 自动生成模板SystemDept
package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// SystemDept 结构体
type SystemDept struct {
      global.GVA_MODEL
      ParentId  *int `json:"parentId" form:"parentId" gorm:"column:parent_id;comment:父ID;size:20;"`
      Level  string `json:"level" form:"level" gorm:"column:level;comment:组级集合;size:500;"`
      Name  string `json:"name" form:"name" gorm:"column:name;comment:部门名称;size:30;"`
      Leader  string `json:"leader" form:"leader" gorm:"column:leader;comment:负责人;size:20;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:联系电话;size:11;"`
      Status  *int `json:"status" form:"status" gorm:"column:status;comment:状态 (1正常 2停用);"`
      Sort  *int `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`
      CreatedBy  *int `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:19;"`
      UpdatedBy  *int `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:19;"`
      Remark  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
}


// TableName SystemDept 表名
func (SystemDept) TableName() string {
  return "system_dept"
}

