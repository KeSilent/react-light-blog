/*
 * @Author: Yang
 * @Date: 2025-04-15 18:05:57
 * @Description: 部门操作类
 */
package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/dal/common/response"
	req "github.com/kesilent/react-light-blog/dal/request"
	"github.com/kesilent/react-light-blog/global"
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
