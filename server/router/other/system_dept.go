package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SystemDeptRouter struct {
}

// InitSystemDeptRouter 初始化 SystemDept 路由信息
func (s *SystemDeptRouter) InitSystemDeptRouter(Router *gin.RouterGroup) {
	systemDeptRouter := Router.Group("systemDept").Use(middleware.OperationRecord())
	systemDeptRouterWithoutRecord := Router.Group("systemDept")
	var systemDeptApi = v1.ApiGroupApp.OtherApiGroup.SystemDeptApi
	{
		systemDeptRouter.POST("createSystemDept", systemDeptApi.CreateSystemDept)             // 新建SystemDept
		systemDeptRouter.DELETE("deleteSystemDept", systemDeptApi.DeleteSystemDept)           // 删除SystemDept
		systemDeptRouter.DELETE("deleteSystemDeptByIds", systemDeptApi.DeleteSystemDeptByIds) // 批量删除SystemDept
		systemDeptRouter.PUT("updateSystemDept", systemDeptApi.UpdateSystemDept)              // 更新SystemDept
	}
	{
		systemDeptRouterWithoutRecord.GET("findSystemDept", systemDeptApi.FindSystemDept)       // 根据ID获取SystemDept
		systemDeptRouterWithoutRecord.GET("getSystemDeptList", systemDeptApi.GetSystemDeptList) // 获取SystemDept列表
		systemDeptRouterWithoutRecord.GET("getSystemDeptTree", systemDeptApi.GetSystemDeptTree) // 获取SystemDept列表
	}
}
