package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SettingGenerateTablesRouter struct {
}

// InitSettingGenerateTablesRouter 初始化 SettingGenerateTables 路由信息
func (s *SettingGenerateTablesRouter) InitSettingGenerateTablesRouter(Router *gin.RouterGroup) {
	settingGenerateTablesRouter := Router.Group("settingGenerateTables").Use(middleware.OperationRecord())
	settingGenerateTablesRouterWithoutRecord := Router.Group("settingGenerateTables")
	var settingGenerateTablesApi = v1.ApiGroupApp.OtherApiGroup.SettingGenerateTablesApi
	{
		settingGenerateTablesRouter.POST("createSettingGenerateTables", settingGenerateTablesApi.CreateSettingGenerateTables)             // 新建SettingGenerateTables
		settingGenerateTablesRouter.DELETE("deleteSettingGenerateTables", settingGenerateTablesApi.DeleteSettingGenerateTables)           // 删除SettingGenerateTables
		settingGenerateTablesRouter.DELETE("deleteSettingGenerateTablesByIds", settingGenerateTablesApi.DeleteSettingGenerateTablesByIds) // 批量删除SettingGenerateTables
		settingGenerateTablesRouter.PUT("updateSettingGenerateTables", settingGenerateTablesApi.UpdateSettingGenerateTables)              // 更新SettingGenerateTables
	}
	{
		settingGenerateTablesRouterWithoutRecord.GET("findSettingGenerateTables", settingGenerateTablesApi.FindSettingGenerateTables)        // 根据ID获取SettingGenerateTables
		settingGenerateTablesRouterWithoutRecord.POST("getSettingGenerateTablesList", settingGenerateTablesApi.GetSettingGenerateTablesList) // 获取SettingGenerateTables列表
	}
}
