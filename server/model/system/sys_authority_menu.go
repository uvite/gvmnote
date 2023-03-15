package system

type SysMenu struct {
	SysBaseMenu
	MenuId      int64                  `json:"menuId" gorm:"comment:菜单ID"`
	AuthorityId uint                   `json:"-" gorm:"comment:角色ID"`
	Children    []SysMenu              `json:"children" gorm:"-"`
	Parameters  []SysBaseMenuParameter `json:"parameters" gorm:"foreignKey:SysBaseMenuID;references:MenuId"`
	Btns        map[string]uint        `json:"btns" gorm:"-"`
}

type SysAuthorityMenu struct {
	MenuId      string `json:"menu_id" gorm:"comment:菜单ID;column:menu_id"`
	AuthorityId string `json:"role_id" gorm:"comment:角色ID;column:role_id"`
}

func (s SysAuthorityMenu) TableName() string {
	return "system_role_menu"
}
