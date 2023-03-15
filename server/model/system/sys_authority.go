package system

import (
	"time"
)

type SysAuthority struct {
	CreatedAt     time.Time  // 创建时间
	UpdatedAt     time.Time  // 更新时间
	DeletedAt     *time.Time `sql:"index"`
	AuthorityId   uint       `json:"id" gorm:"primarykey"`          // 主键ID
	AuthorityName string     `json:"name" gorm:"comment:角色名"`       // 角色名
	Code          string     `json:"code" gorm:"comment:角色名"`       // 角色名
	DataScope     int        `json:"data_scope" gorm:"comment:角色名"` // 角色名
	Status        int        `json:"status" gorm:"comment:类型"`      // 是否在列表隐藏
	Score         int        `json:"score" gorm:"comment:类型"`       // 是否在列表隐藏
	Remark        string     `json:"remark" gorm:"comment:类型"`      // 是否在列表隐藏

	ParentId        *uint           `json:"parentId" gorm:"comment:父角色ID"` // 父角色ID
	DataAuthorityId []*SysAuthority `json:"dataAuthorityId" gorm:"many2many:sys_data_authority_id;"`
	Children        []SysAuthority  `json:"children" gorm:"-"`
	SysBaseMenus    []SysBaseMenu   `json:"menus" gorm:"many2many:sys_authority_menus;"`
	Users           []SysUser       `json:"-" gorm:"many2many:sys_user_authority;"`
	DefaultRouter   string          `json:"defaultRouter" gorm:"comment:默认菜单;default:dashboard"` // 默认菜单(默认dashboard)
}

func (SysAuthority) TableName() string {
	return "system_role_1"
}
