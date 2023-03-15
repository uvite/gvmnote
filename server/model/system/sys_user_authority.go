package system

// SysUserAuthority 是 sysUser 和 sysAuthority 的连接表
type SysUserAuthority struct {
	SysUserId               uint `gorm:"column:user_id"`
	SysAuthorityAuthorityId uint `gorm:"column:role_id"`
}

func (s *SysUserAuthority) TableName() string {
	return "system_user_role"
}
