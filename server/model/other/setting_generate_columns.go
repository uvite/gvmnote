// 自动生成模板SettingGenerateColumns
package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SettingGenerateColumns 结构体
type SettingGenerateColumns struct {
	global.GVA_MODEL
	SettingGenerateTablesID int    `json:"table_id" form:"table_id" gorm:"column:table_id;comment:所属表ID"` // 关联标记
	ColumnName              string `json:"column_name" form:"column_name" gorm:"column:column_name;comment:字段名称;size:200;"`
	ColumnComment           string `json:"column_comment" form:"column_comment" gorm:"column:column_comment;comment:字段注释;size:255;"`
	ColumnType              string `json:"column_type" form:"column_type" gorm:"column:column_type;comment:字段类型;size:50;"`
	IsPk                    *int   `json:"is_pk" form:"is_pk" gorm:"column:is_pk;comment:1 非主键 2 主键;"`
	IsRequired              *int   `json:"is_required" form:"is_required" gorm:"column:is_required;comment:1 非必填 2 必填;"`
	IsInsert                *int   `json:"is_insert" form:"is_insert" gorm:"column:is_insert;comment:1 非插入字段 2 插入字段;"`
	IsEdit                  *int   `json:"is_edit" form:"is_edit" gorm:"column:is_edit;comment:1 非编辑字段 2 编辑字段;"`
	IsList                  *int   `json:"is_list" form:"is_list" gorm:"column:is_list;comment:1 非列表显示字段 2 列表显示字段;"`
	IsQuery                 *int   `json:"is_query" form:"is_query" gorm:"column:is_query;comment:1 非查询字段 2 查询字段;"`
	IsSort                  *int   `json:"is_sort" form:"is_sort" gorm:"column:is_sort;comment:1 不排序 2 排序字段;"`
	QueryType               string `json:"query_type" form:"query_type" gorm:"column:query_type;comment:查询方式 eq 等于, neq 不等于, gt 大于, lt 小于, like 范围;size:100;"`
	ViewType                string `json:"view_type" form:"view_type" gorm:"column:view_type;comment:页面控件，text, textarea, password, select, checkbox, radio, date, upload, ma-upload（封装的上传控件）;size:100;"`
	DictType                string `json:"dict_type" form:"dict_type" gorm:"column:dict_type;comment:字典类型;size:200;"`
	AllowRoles              string `json:"allow_roles" form:"allow_roles" gorm:"column:allow_roles;comment:允许查看该字段的角色;size:255;"`
	Options                 string `json:"options" form:"options" gorm:"column:options;comment:字段其他设置;size:1000;"`
	Sort                    *bool  `json:"sort" form:"sort" gorm:"column:sort;comment:排序;"`
	CreatedBy               *int   `json:"created_by" form:"created_by" gorm:"column:created_by;comment:创建者;size:19;"`
	UpdatedBy               *int   `json:"updated_by" form:"updated_by" gorm:"column:updated_by;comment:更新者;size:19;"`
	Remark                  string `json:"remark" form:"remark" gorm:"column:remark;comment:备注;size:255;"`
}

// TableName SettingGenerateColumns 表名
func (SettingGenerateColumns) TableName() string {
	return "setting_generate_columns"
}
