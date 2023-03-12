package other

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/other"
	otherReq "github.com/flipped-aurora/gin-vue-admin/server/model/other/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type SettingGenerateColumnsApi struct {
}

var settingGenerateColumnsService = service.ServiceGroupApp.OtherServiceGroup.SettingGenerateColumnsService

// CreateSettingGenerateColumns 创建SettingGenerateColumns
// @Tags SettingGenerateColumns
// @Summary 创建SettingGenerateColumns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SettingGenerateColumns true "创建SettingGenerateColumns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /settingGenerateColumns/createSettingGenerateColumns [post]
func (settingGenerateColumnsApi *SettingGenerateColumnsApi) CreateSettingGenerateColumns(c *gin.Context) {
	var settingGenerateColumns other.SettingGenerateColumns
	err := c.ShouldBindJSON(&settingGenerateColumns)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingGenerateColumnsService.CreateSettingGenerateColumns(settingGenerateColumns); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSettingGenerateColumns 删除SettingGenerateColumns
// @Tags SettingGenerateColumns
// @Summary 删除SettingGenerateColumns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SettingGenerateColumns true "删除SettingGenerateColumns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /settingGenerateColumns/deleteSettingGenerateColumns [delete]
func (settingGenerateColumnsApi *SettingGenerateColumnsApi) DeleteSettingGenerateColumns(c *gin.Context) {
	var settingGenerateColumns other.SettingGenerateColumns
	err := c.ShouldBindJSON(&settingGenerateColumns)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingGenerateColumnsService.DeleteSettingGenerateColumns(settingGenerateColumns); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSettingGenerateColumnsByIds 批量删除SettingGenerateColumns
// @Tags SettingGenerateColumns
// @Summary 批量删除SettingGenerateColumns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SettingGenerateColumns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /settingGenerateColumns/deleteSettingGenerateColumnsByIds [delete]
func (settingGenerateColumnsApi *SettingGenerateColumnsApi) DeleteSettingGenerateColumnsByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingGenerateColumnsService.DeleteSettingGenerateColumnsByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSettingGenerateColumns 更新SettingGenerateColumns
// @Tags SettingGenerateColumns
// @Summary 更新SettingGenerateColumns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SettingGenerateColumns true "更新SettingGenerateColumns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /settingGenerateColumns/updateSettingGenerateColumns [put]
func (settingGenerateColumnsApi *SettingGenerateColumnsApi) UpdateSettingGenerateColumns(c *gin.Context) {
	var settingGenerateColumns other.SettingGenerateColumns
	err := c.ShouldBindJSON(&settingGenerateColumns)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingGenerateColumnsService.UpdateSettingGenerateColumns(settingGenerateColumns); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSettingGenerateColumns 用id查询SettingGenerateColumns
// @Tags SettingGenerateColumns
// @Summary 用id查询SettingGenerateColumns
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query other.SettingGenerateColumns true "用id查询SettingGenerateColumns"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /settingGenerateColumns/findSettingGenerateColumns [get]
func (settingGenerateColumnsApi *SettingGenerateColumnsApi) FindSettingGenerateColumns(c *gin.Context) {
	var settingGenerateColumns other.SettingGenerateColumns
	err := c.ShouldBindQuery(&settingGenerateColumns)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resettingGenerateColumns, err := settingGenerateColumnsService.GetSettingGenerateColumns(settingGenerateColumns.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resettingGenerateColumns": resettingGenerateColumns}, c)
	}
}

// GetSettingGenerateColumnsList 分页获取SettingGenerateColumns列表
// @Tags SettingGenerateColumns
// @Summary 分页获取SettingGenerateColumns列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query otherReq.SettingGenerateColumnsSearch true "分页获取SettingGenerateColumns列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /settingGenerateColumns/getSettingGenerateColumnsList [get]
func (settingGenerateColumnsApi *SettingGenerateColumnsApi) GetSettingGenerateColumnsList(c *gin.Context) {
	var pageInfo otherReq.SettingGenerateColumnsSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := settingGenerateColumnsService.GetSettingGenerateColumnsInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (settingGenerateColumnsApi *SettingGenerateColumnsApi) GetColumnsByTableId(c *gin.Context) {

	var json other.SettingGenerateColumns
	err := c.ShouldBindQuery(&json)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	db := global.GVA_DB.Model(&other.SettingGenerateColumns{})
	var settingGenerateColumns []other.SettingGenerateColumns

	db = db.Where("table_id = ? ", json.SettingGenerateTablesID)

	err = db.Find(&settingGenerateColumns).Error

	response.OkWithDetailed(settingGenerateColumns, "获取成功", c)
}
