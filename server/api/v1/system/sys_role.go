/*
 * @Author: Yang
 * @Date: 2025-03-19 17:03:23
 * @Description: 角色权限API
 */
package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/dal/common/response"
	"github.com/kesilent/react-light-blog/dal/model"
	systemReq "github.com/kesilent/react-light-blog/dal/request"
	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/utils"
	"go.uber.org/zap"
)

type RoleApi struct{}

/**
 * @Author: kesilent
 * @Description: 保存角色
 **/
func (r *RoleApi) SaveRole(c *gin.Context) {
	var role model.SysRole
	err := c.ShouldBindJSON(&role)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = utils.Verify(role, utils.SysRoleVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if role.ID == 0 {
		role.ID = model.SnowflakeID(utils.GenID(1))
	}
	err = roleService.SaveRole(&role)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}
	response.Ok(c)
}

/**
 * @Author: kesilent
 * @Description: 删除角色
 **/
func (r *RoleApi) DeleteRole(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		response.FailWithMessage("id不能为空", c)
		return
	}
	idInt, _ := utils.StrToInt64(id)
	resultInfo, err := roleService.DeleteRole(idInt)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
	}

	if resultInfo.RowsAffected > 0 {
		response.OkWithMessage("删除"+fmt.Sprintf("%d", resultInfo.RowsAffected)+"条", c)

	} else {
		response.FailWithMessage("删除失败", c)
	}
}

/**
 * @Author: kesilent
 * @Description: 获取角色列表
 **/
func (r *RoleApi) GetRoleList(c *gin.Context) {
	var pageInfo systemReq.GetRoleListReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	list, total, err := roleService.GetRoleList(pageInfo)
	if err != nil {
		global.RLB_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		Data:     list,
		Total:    total,
		Page:     pageInfo.Current,
		PageSize: pageInfo.PageSize,
		Success:  true,
	}, "获取成功", c)
}

/**
 * @Author: kesilent
 * @Description: 获取角色菜单
 **/
func (r *RoleApi) GetRoleMenus(c *gin.Context) {
	roleId := c.Query("roleId")
	if roleId == "" {
		response.FailWithMessage("角色ID不能为空", c)
		return
	}
	roleIdInt, _ := utils.StrToInt64(roleId)
	menus, err := roleService.GetRoleMenus(roleIdInt)
	if err != nil {
		global.RLB_LOG.Error("获取角色菜单失败!", zap.Error(err))
		response.FailWithMessage("获取角色菜单失败", c)
		return
	}
	response.OkWithDetailed(menus, "获取角色菜单成功", c)
}

/**
 * @Author: kesilent
 * @Description: 添加角色菜单
 **/
func (r *RoleApi) AddRoleMenu(c *gin.Context) {
	var roleMenu systemReq.RoleMenuReq
	err := c.BindJSON(&roleMenu)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}

	err = utils.Verify(roleMenu, utils.RoleMenuReqVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = roleService.AddRoleMenus(roleMenu)
	if err != nil {
		global.RLB_LOG.Error("添加角色菜单失败!", zap.Error(err))
		response.FailWithMessage("添加角色菜单失败", c)
		return
	}
	response.OkWithMessage("添加角色菜单成功", c)
}
