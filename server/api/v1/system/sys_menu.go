/*
 * @Author: KeSilent kesilent@live.com
 * @Date: 2025-03-19 19:04:24
 * @Description:菜单操作类
 * @LastEditors: Please set LastEditors
 * @LastEditTime: 2025-04-10 16:16:36
 * @FilePath: /server/api/v1/system/sys_menu.go
 *
 * Copyright (c) 2025 by ${git_name_email}, All Rights Reserved.
 */
package system

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/dal/common/response"
	"github.com/kesilent/react-light-blog/dal/model"
	req "github.com/kesilent/react-light-blog/dal/request"
	"github.com/kesilent/react-light-blog/global"
	"github.com/kesilent/react-light-blog/utils"
	"go.uber.org/zap"
)

type MenuApi struct{}

/**
 * @description: 登录获取菜单
 * @param {*gin.Context} c
 * @return {*}
 */
func (m *MenuApi) Menus(c *gin.Context) {
	user, _ := userService.GetUserInfo(utils.GetUserUuid(c))
	if len(user.Role) == 0 {
		response.FailWithMessage("未设置任何角色", c)
	}

	//获取第一个角色
	menus, err := menuService.GetRoleMenuList(user.Role[0].UUID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(menus, c)
		return
	}
}

/***
 * @Author: kesilent
 * @Description: 获取菜单列表
 **/
func (m *MenuApi) GetMenuByKey(c *gin.Context) {
	keyWord := c.Query("keyWord")
	//获取列表
	menus, err := menuService.GetMenuByKey(keyWord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	} else {
		response.OkWithData(menus, c)
		return
	}
}

/**
 * @Author: Yang
 * @description: 获取菜单分页列表
 * @param {*gin.Context} c
 * @return {*}
 */
func (m *MenuApi) GetMenuListByPage(c *gin.Context) {
	var pageInfo req.GetMenuListReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := menuService.GetMenuListByPage(pageInfo)
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
 * @Author: Yang
 * @description: 保存菜单
 * @param {*gin.Context} c
 * @return {*}
 */
func (m *MenuApi) SaveBaseMenu(c *gin.Context) {
	var menu model.SysBaseMenu
	err := c.ShouldBindJSON(&menu)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := menuService.SaveBaseMenu(menu); err != nil {
		global.RLB_LOG.Error("添加失败!", zap.Error(err))
		response.FailWithMessage("添加失败", c)
		return
	} else {
		response.OkWithMessage("添加成功", c)
		return
	}
}

/**
 * @Author: Yang
 * @description: 删除菜单
 * @param {*gin.Context} c
 * @return {*}
 */
func (m *MenuApi) DeleteMenu(c *gin.Context) {
	uuid := c.Query("id")
	resultInfo, err := menuService.DeleteMenu(uuid)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resultInfo.RowsAffected > 0 {
		response.OkWithMessage("删除"+fmt.Sprintf("%d", resultInfo.RowsAffected)+"条", c)

	} else {
		response.FailWithMessage("删除失败", c)
	}
}
