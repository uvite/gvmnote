package system

type SysRoleMenu struct {
	MenuId string `json:"menu_id" gorm:"comment:菜单ID;column:sys_base_menu_id"`
	RoleId string `json:"role_id" gorm:"comment:角色ID;column:system_role_id"`
}

func (s SysRoleMenu) TableName() string {
	return "system_role_menu"
}
