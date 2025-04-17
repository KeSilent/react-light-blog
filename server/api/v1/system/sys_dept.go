/*
 * @Author: Yang
 * @Date: 2025-04-15 18:05:57
 * @Description: 部门操作类
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

type DeptApi struct {
}

/**
 * @Author: Yang
 * @description: 获取部门分页列表
 * @param {*gin.Context} c
 * @return {*}
 */
func (deptApi *DeptApi) GetList(c *gin.Context) {
	var pageInfo req.GetDeptListReq
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	list, total, err := deptService.GetListByPage(pageInfo)
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
 * @description: 获取下拉菜单部门列表
 * @param {*gin.Context} c
 * @return {*}
 */
func (deptApi *DeptApi) GetListByTreeSelect(c *gin.Context) {
	key := c.Query("key")
	list, err := deptService.GetListByTreeSelect(key)
	if err != nil {
		global.RLB_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(list, "获取成功", c)
}

/**
 * @Author: Yang
 * @description: 保存部门
 * @param {*gin.Context} c
 * @return {*}
 */
func (deptApi *DeptApi) SaveDept(c *gin.Context) {
	var dept model.SysDept
	err := c.ShouldBindJSON(&dept)
	if err != nil {
		response.FailWithMessage("参数错误", c)
		return
	}
	if err := deptService.SaveDept(dept); err != nil {
		global.RLB_LOG.Error("保存失败!", zap.Error(err))
		response.FailWithMessage("保存失败", c)
		return
	}
	response.OkWithMessage("保存成功", c)
}

/**
 * @Author: Yang
 * @description: 删除部门
 * @param {*gin.Context} c
 * @return {*}
 */
func (deptApi *DeptApi) DeleteDept(c *gin.Context) {
	id := c.Query("id")
	idInt, _ := utils.StrToInt64(id)
	resultInfo, err := deptService.DeleteDept(idInt)
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
