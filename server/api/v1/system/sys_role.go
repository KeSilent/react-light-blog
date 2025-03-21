package system

import (
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
		response.FailWithMessage("参数错误", c)
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

/**
 * @Author: kesilent
 * @Description: 添加角色菜单
 **/
func (r *RoleApi) AddRoleMenu(c *gin.Context) {
	var roleMenu []*model.SysRoleMenu
	err := c.BindJSON(&roleMenu)
	if err != nil {
		response.FailWithMessage("参数错误", c)
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
