// 自动生成模板SettingGenerateTables
package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SettingGenerateTables 结构体
type SettingGenerateTables struct {
	global.GVA_MODEL
	Tablealias    string `json:"table_name" form:"table_name" gorm:"column:table_name;comment:表名称;size:200;"`
	TableComment  string `json:"table_comment" form:"table_comment" gorm:"column:table_comment;comment:表注释;size:500;"`
	ModuleName    string `json:"module_name" form:"module_name" gorm:"column:module_name;comment:所属模块;size:100;"`
	Namespace     string `json:"namespace" form:"namespace" gorm:"column:namespace;comment:命名空间;size:255;"`
	MenuName      string `json:"menu_name" form:"menu_name" gorm:"column:menu_name;comment:生成菜单名;size:100;"`
	BelongMenuId  *int   `json:"belong_menu_id" form:"belong_menu_id" gorm:"column:belong_menu_id;comment:所属菜单;size:19;"`
	PackageName   string `json:"package_name" form:"package_name" gorm:"column:package_name;comment:控制器包名;size:100;"`
	Type          string `json:"type" form:"type" gorm:"column:type;comment:生成类型，single 单表CRUD，tree 树表CRUD，parent_sub父子表CRUD;size:100;"`
	GenerateType  *int   `json:"generate_type" form:"generate_type" gorm:"column:generate_type;comment:1 压缩包下载 2 生成到模块;"`
	GenerateMenus string `json:"generate_menus" form:"generate_menus" gorm:"column:generate_menus;comment:生成菜单列表;size:255;"`
	BuildMenu     *int   `json:"build_menu" form:"build_menu" gorm:"column:build_menu;comment:是否构建菜单;"`
	ComponentType *int   `json:"component_type" form:"component_type" gorm:"column:component_type;comment:组件显示方式;"`
	Options       string `json:"options" form:"options" gorm:"column:options;comment:其他业务选项;size:1500;"`
	CreatedBy     *int   `json:"created_by" form:"created_by" gorm:"column:created_by;comment:创建者;size:19;"`
	UpdatedBy     *int   `json:"updated_by" form:"updated_by" gorm:"column:updated_by;comment:更新者;size:19;"`
	Remark        string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
	SettingGenerateColumns []SettingGenerateColumns `json:"settingGenerateColumns" form:"settingGenerateColumns"`

}

// TableName SettingGenerateTables 表名
func (SettingGenerateTables) TableName() string {
	return "setting_generate_tables"
}
