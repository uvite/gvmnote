package system

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/satori/go.uuid"
)

type SysUser struct {
	global.GVA_MODEL
	UUID        uuid.UUID      `json:"uuid" gorm:"index;comment:用户UUID"`
	Username    string         `json:"username" gorm:"index;comment:用户登录名"`                                               // 用户登录名
	Password    string         `json:"-"  gorm:"comment:用户登录密码"`                                                          // 用户登录密码
	NickName    string         `json:"nickname" gorm:"default:系统用户;comment:用户昵称"`                                         // 用户昵称
	SideMode    string         `json:"sideMode" gorm:"default:dark;comment:用户侧边主题"`                                       // 用户侧边主题
	HeaderImg   string         `json:"avatar" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:用户头像"` // 用户头像
	BaseColor   string         `json:"baseColor" gorm:"default:#fff;comment:基础颜色"`                                        // 基础颜色
	ActiveColor string         `json:"activeColor" gorm:"default:#1890ff;comment:活跃颜色"`                                   // 活跃颜色
	AuthorityId uint           `json:"authorityId" gorm:"default:888;comment:用户角色ID"`                                     // 用户角色ID
	Authority   SysAuthority   `json:"authority" gorm:"foreignKey:AuthorityId;references:AuthorityId;comment:用户角色"`
	Authorities []SysAuthority `json:"authorities" gorm:"many2many:sys_user_authority;"`
	Phone       string         `json:"phone"  gorm:"comment:用户手机号"`                          // 用户手机号
	Email       string         `json:"email"  gorm:"comment:用户邮箱"`                           // 用户邮箱
	Signed      string         `json:"signed"  gorm:"comment:用户邮箱"`                          // 用户邮箱
	Dashboard   string         `json:"dashboard"  gorm:"comment:用户邮箱"`                       // 用户邮箱
	Enable      int            `json:"enable" gorm:"default:1;comment:用户是否被冻结 1正常 2冻结"`      //用户是否被冻结 1正常 2冻结
	Status      int            `json:"status" form:"status" gorm:"column:status;comment:状态"` // 状态

	BackendSetting BackendSetting `json:"backend_setting" gorm:"type:bytes;serializer:gob"` // 用户邮箱

}

type BackendSetting struct {
	Mode         string `json:"mode"`
	Tag          bool   `json:"tag"`
	MenuCollapse bool   `json:"menuCollapse"`
	MenuWidth    int    `json:"menuWidth"`
	Layout       string `json:"layout"`
	Skin         string `json:"skin"`
	I18N         bool   `json:"i18n"`
	Language     string `json:"language"`
	Animation    string `json:"animation"`
	Color        string `json:"color"`
}

func (SysUser) TableName() string {
	return "system_user"
}
