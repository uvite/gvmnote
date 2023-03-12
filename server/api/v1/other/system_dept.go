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
	"math"
)

type SystemDeptApi struct {
}

var systemDeptService = service.ServiceGroupApp.OtherServiceGroup.SystemDeptService

// CreateSystemDept 创建SystemDept
// @Tags SystemDept
// @Summary 创建SystemDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SystemDept true "创建SystemDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemDept/createSystemDept [post]
func (systemDeptApi *SystemDeptApi) CreateSystemDept(c *gin.Context) {
	var systemDept other.SystemDept
	err := c.ShouldBindJSON(&systemDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := systemDeptService.CreateSystemDept(systemDept); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSystemDept 删除SystemDept
// @Tags SystemDept
// @Summary 删除SystemDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SystemDept true "删除SystemDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /systemDept/deleteSystemDept [delete]
func (systemDeptApi *SystemDeptApi) DeleteSystemDept(c *gin.Context) {
	var systemDept other.SystemDept
	err := c.ShouldBindJSON(&systemDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := systemDeptService.DeleteSystemDept(systemDept); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSystemDeptByIds 批量删除SystemDept
// @Tags SystemDept
// @Summary 批量删除SystemDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SystemDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /systemDept/deleteSystemDeptByIds [delete]
func (systemDeptApi *SystemDeptApi) DeleteSystemDeptByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := systemDeptService.DeleteSystemDeptByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSystemDept 更新SystemDept
// @Tags SystemDept
// @Summary 更新SystemDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SystemDept true "更新SystemDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /systemDept/updateSystemDept [put]
func (systemDeptApi *SystemDeptApi) UpdateSystemDept(c *gin.Context) {
	var systemDept other.SystemDept
	err := c.ShouldBindJSON(&systemDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := systemDeptService.UpdateSystemDept(systemDept); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSystemDept 用id查询SystemDept
// @Tags SystemDept
// @Summary 用id查询SystemDept
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query other.SystemDept true "用id查询SystemDept"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /systemDept/findSystemDept [get]
func (systemDeptApi *SystemDeptApi) FindSystemDept(c *gin.Context) {
	var systemDept other.SystemDept
	err := c.ShouldBindQuery(&systemDept)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resystemDept, err := systemDeptService.GetSystemDept(systemDept.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resystemDept": resystemDept}, c)
	}
}

// GetSystemDeptList 分页获取SystemDept列表
// @Tags SystemDept
// @Summary 分页获取SystemDept列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query otherReq.SystemDeptSearch true "分页获取SystemDept列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemDept/getSystemDeptList [get]
func (systemDeptApi *SystemDeptApi) GetSystemDeptList(c *gin.Context) {
	var pageInfo otherReq.SystemDeptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := systemDeptService.GetSystemDeptInfoList(pageInfo); err != nil {
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

func (systemDeptApi *SystemDeptApi) GetSystemDeptTree(c *gin.Context) {
	var pageInfo otherReq.SystemDeptSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, _, err := systemDeptService.GetSystemDeptTree(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {

		response.OkWithDetailed(list, "获取成功", c)
	}
}
