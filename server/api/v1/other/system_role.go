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

type SystemRoleApi struct {
}

var systemRoleService = service.ServiceGroupApp.OtherServiceGroup.SystemRoleService


// CreateSystemRole 创建SystemRole
// @Tags SystemRole
// @Summary 创建SystemRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SystemRole true "创建SystemRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemRole/createSystemRole [post]
func (systemRoleApi *SystemRoleApi) CreateSystemRole(c *gin.Context) {
	var systemRole other.SystemRole
	err := c.ShouldBindJSON(&systemRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := systemRoleService.CreateSystemRole(systemRole); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteSystemRole 删除SystemRole
// @Tags SystemRole
// @Summary 删除SystemRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SystemRole true "删除SystemRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /systemRole/deleteSystemRole [delete]
func (systemRoleApi *SystemRoleApi) DeleteSystemRole(c *gin.Context) {
	var systemRole other.SystemRole
	err := c.ShouldBindJSON(&systemRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := systemRoleService.DeleteSystemRole(systemRole); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteSystemRoleByIds 批量删除SystemRole
// @Tags SystemRole
// @Summary 批量删除SystemRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SystemRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /systemRole/deleteSystemRoleByIds [delete]
func (systemRoleApi *SystemRoleApi) DeleteSystemRoleByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := systemRoleService.DeleteSystemRoleByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateSystemRole 更新SystemRole
// @Tags SystemRole
// @Summary 更新SystemRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body other.SystemRole true "更新SystemRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /systemRole/updateSystemRole [put]
func (systemRoleApi *SystemRoleApi) UpdateSystemRole(c *gin.Context) {
	var systemRole other.SystemRole
	err := c.ShouldBindJSON(&systemRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := systemRoleService.UpdateSystemRole(systemRole); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindSystemRole 用id查询SystemRole
// @Tags SystemRole
// @Summary 用id查询SystemRole
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query other.SystemRole true "用id查询SystemRole"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /systemRole/findSystemRole [get]
func (systemRoleApi *SystemRoleApi) FindSystemRole(c *gin.Context) {
	var systemRole other.SystemRole
	err := c.ShouldBindQuery(&systemRole)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resystemRole, err := systemRoleService.GetSystemRole(systemRole.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resystemRole": resystemRole}, c)
	}
}

// GetSystemRoleList 分页获取SystemRole列表
// @Tags SystemRole
// @Summary 分页获取SystemRole列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query otherReq.SystemRoleSearch true "分页获取SystemRole列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /systemRole/getSystemRoleList [get]
func (systemRoleApi *SystemRoleApi) GetSystemRoleList(c *gin.Context) {
	var pageInfo otherReq.SystemRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := systemRoleService.GetSystemRoleInfoList(pageInfo); err != nil {
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
func (systemRoleApi *SystemRoleApi) GetSystemRoleAll(c *gin.Context) {
	var pageInfo otherReq.SystemRoleSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, _, err := systemRoleService.GetSystemRoleInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(list ,"获取成功", c)
    }
}
