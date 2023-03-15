package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SystemRoleRouter struct {
}

// InitSystemRoleRouter 初始化 SystemRole 路由信息
func (s *SystemRoleRouter) InitSystemRoleRouter(Router *gin.RouterGroup) {
	systemRoleRouter := Router.Group("systemRole").Use(middleware.OperationRecord())
	systemRoleRouterWithoutRecord := Router.Group("systemRole")
	var systemRoleApi = v1.ApiGroupApp.OtherApiGroup.SystemRoleApi
	authorityMenuApi := v1.ApiGroupApp.SystemApiGroup.AuthorityMenuApi

	{
		systemRoleRouter.POST("createSystemRole", systemRoleApi.CreateSystemRole)             // 新建SystemRole
		systemRoleRouter.DELETE("deleteSystemRole", systemRoleApi.DeleteSystemRole)           // 删除SystemRole
		systemRoleRouter.DELETE("deleteSystemRoleByIds", systemRoleApi.DeleteSystemRoleByIds) // 批量删除SystemRole
		systemRoleRouter.PUT("updateSystemRole", systemRoleApi.UpdateSystemRole)              // 更新SystemRole
		systemRoleRouter.POST("addMenuAuthority", authorityMenuApi.AddMenuAuthority)          //	增加menu和角色关联关系

	}
	{
		systemRoleRouterWithoutRecord.POST("getMenuAuthority", authorityMenuApi.GetMenuAuthority) // 获取指定角色menu

		systemRoleRouterWithoutRecord.GET("findSystemRole", systemRoleApi.FindSystemRole)       // 根据ID获取SystemRole
		systemRoleRouterWithoutRecord.GET("getSystemRoleList", systemRoleApi.GetSystemRoleList) // 获取SystemRole列表
		systemRoleRouterWithoutRecord.GET("getSystemRoleAll", systemRoleApi.GetSystemRoleAll)   // 获取SystemRole列表
	}
}
