package other

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/other"
	otherReq "github.com/flipped-aurora/gin-vue-admin/server/model/other/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"math"
)

type SettingGenerateTablesApi struct {
}

var settingGenerateTablesService = service.ServiceGroupApp.OtherServiceGroup.SettingGenerateTablesService

type PostTables struct {
	TableName    string `json:"table_name" form:"table_name"`       // 页码
	TableComment string `json:"table_comment" form:"table_comment"` // 每页大小
}
type GenerateTables struct {
	BusinessDB string       `json:"businessDB" form:"businessDB"` // 页码
	Database   string       `json:"database" form:"database"`     // 每页大小
	Tables     []PostTables `json:"tables" form:"tables"`         //关键字
}

func (settingGenerateTablesApi *SettingGenerateTablesApi) CreateSettingGenerateTables(c *gin.Context) {
	var generateTables GenerateTables
	err := c.ShouldBindJSON(&generateTables)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	fmt.Println(generateTables)
	for key, val := range generateTables.Tables {
		fmt.Println(key, val)

		generateTable := other.SettingGenerateTables{
			TableComment: val.TableComment,
			Tablealias:   val.TableName,
		}
		columns, err := system.AutoCodeServiceApp.Database(generateTables.BusinessDB).GetColumn(generateTables.BusinessDB, val.TableName, generateTables.Database)
		fmt.Println(err)
		for key, value := range columns {
			fmt.Println(key, value)
			column := other.SettingGenerateColumns{
				ColumnName:    value.ColumnName,
				ColumnComment: value.ColumnComment,
				ColumnType:    value.DataType,
			}
			generateTable.SettingGenerateColumns = append(generateTable.SettingGenerateColumns, column)
		}

		err = global.GVA_DB.Create(&generateTable).Error

	}

	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSettingGenerateTables 删除SettingGenerateTables
// @Tags SettingGenerateTables
// @Summary 删除SettingGenerateTables
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SettingGenerateTables true "删除SettingGenerateTables"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /settingGenerateTables/deleteSettingGenerateTables [delete]
func (settingGenerateTablesApi *SettingGenerateTablesApi) DeleteSettingGenerateTables(c *gin.Context) {
	var settingGenerateTables other.SettingGenerateTables
	err := c.ShouldBindJSON(&settingGenerateTables)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingGenerateTablesService.DeleteSettingGenerateTables(settingGenerateTables); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSettingGenerateTablesByIds 批量删除SettingGenerateTables
// @Tags SettingGenerateTables
// @Summary 批量删除SettingGenerateTables
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SettingGenerateTables"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /settingGenerateTables/deleteSettingGenerateTablesByIds [delete]
func (settingGenerateTablesApi *SettingGenerateTablesApi) DeleteSettingGenerateTablesByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingGenerateTablesService.DeleteSettingGenerateTablesByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSettingGenerateTables 更新SettingGenerateTables
// @Tags SettingGenerateTables
// @Summary 更新SettingGenerateTables
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SettingGenerateTables true "更新SettingGenerateTables"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /settingGenerateTables/updateSettingGenerateTables [put]
func (settingGenerateTablesApi *SettingGenerateTablesApi) UpdateSettingGenerateTables(c *gin.Context) {
	var settingGenerateTables other.SettingGenerateTables
	err := c.ShouldBindJSON(&settingGenerateTables)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := settingGenerateTablesService.UpdateSettingGenerateTables(settingGenerateTables); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSettingGenerateTables 用id查询SettingGenerateTables
// @Tags SettingGenerateTables
// @Summary 用id查询SettingGenerateTables
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query other.SettingGenerateTables true "用id查询SettingGenerateTables"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /settingGenerateTables/findSettingGenerateTables [get]
func (settingGenerateTablesApi *SettingGenerateTablesApi) FindSettingGenerateTables(c *gin.Context) {
	var settingGenerateTables other.SettingGenerateTables
	err := c.ShouldBindQuery(&settingGenerateTables)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resettingGenerateTables, err := settingGenerateTablesService.GetSettingGenerateTables(settingGenerateTables.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(resettingGenerateTables, c)
	}
}

// GetSettingGenerateTablesList 分页获取SettingGenerateTables列表
// @Tags SettingGenerateTables
// @Summary 分页获取SettingGenerateTables列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query otherReq.SettingGenerateTablesSearch true "分页获取SettingGenerateTables列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /settingGenerateTables/getSettingGenerateTablesList [get]
func (settingGenerateTablesApi *SettingGenerateTablesApi) GetSettingGenerateTablesList(c *gin.Context) {
	var pageInfo otherReq.SettingGenerateTablesSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := settingGenerateTablesService.GetSettingGenerateTablesInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {

		var p = response.ResponsePageInfo{
			Total:       total,
			TotalPage:   int64(math.Ceil(float64((total / int64(pageInfo.PageSize))))),
			CurrentPage: int64(pageInfo.Page),
		}
		response.OkWithDetailed(response.ReturnPage{
			List:             list,
			ResponsePageInfo: p,
		}, "获取成功", c)

	}
}
