package system

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/dal/common/response"
	systemReq "github.com/kesilent/react-light-blog/dal/request"
	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/utils"
	"go.uber.org/zap"
)

type RoleApi struct{}

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
	roleIdStr := c.Query("roleId")
	if roleIdStr == "" {
		response.FailWithMessage("参数错误", c)
		return
	}
	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		response.FailWithMessage("角色ID格式错误", c)
		return
	}
	menus, err := roleService.GetRoleMenus(roleId)
	if err != nil {
		global.RLB_LOG.Error("获取角色菜单失败!", zap.Error(err))
		response.FailWithMessage("获取角色菜单失败", c)
		return
	}
	response.OkWithDetailed(menus, "获取角色菜单成功", c)
}
