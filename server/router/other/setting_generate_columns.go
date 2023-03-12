package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SettingGenerateColumnsRouter struct {
}

// InitSettingGenerateColumnsRouter 初始化 SettingGenerateColumns 路由信息
func (s *SettingGenerateColumnsRouter) InitSettingGenerateColumnsRouter(Router *gin.RouterGroup) {
	settingGenerateColumnsRouter := Router.Group("settingGenerateColumns").Use(middleware.OperationRecord())
	settingGenerateColumnsRouterWithoutRecord := Router.Group("settingGenerateColumns")
	var settingGenerateColumnsApi = v1.ApiGroupApp.OtherApiGroup.SettingGenerateColumnsApi
	{
		settingGenerateColumnsRouter.POST("createSettingGenerateColumns", settingGenerateColumnsApi.CreateSettingGenerateColumns)   // 新建SettingGenerateColumns
		settingGenerateColumnsRouter.DELETE("deleteSettingGenerateColumns", settingGenerateColumnsApi.DeleteSettingGenerateColumns) // 删除SettingGenerateColumns
		settingGenerateColumnsRouter.DELETE("deleteSettingGenerateColumnsByIds", settingGenerateColumnsApi.DeleteSettingGenerateColumnsByIds) // 批量删除SettingGenerateColumns
		settingGenerateColumnsRouter.PUT("updateSettingGenerateColumns", settingGenerateColumnsApi.UpdateSettingGenerateColumns)    // 更新SettingGenerateColumns
	}
	{
		settingGenerateColumnsRouterWithoutRecord.GET("findSettingGenerateColumns", settingGenerateColumnsApi.FindSettingGenerateColumns)        // 根据ID获取SettingGenerateColumns
		settingGenerateColumnsRouterWithoutRecord.GET("getSettingGenerateColumnsList", settingGenerateColumnsApi.GetSettingGenerateColumnsList)  // 获取SettingGenerateColumns列表
		settingGenerateColumnsRouterWithoutRecord.GET("getColumnsByTableId", settingGenerateColumnsApi.GetColumnsByTableId)        // 根据ID获取SettingGenerateColumns

	}
}
