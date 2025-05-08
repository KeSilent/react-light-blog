package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/dal/common/response"
	"github.com/kesilent/react-light-blog/global"
	"go.uber.org/zap"
)

type AutoCodeApi struct {
}

/**
 * @Author: Yang
 * @description: 获取所有的表
 * @param {*gin.Context} c
 * @return {*}
 */
func (a *AutoCodeApi) GetAllTableName(c *gin.Context) {
	list, err := autoCodeService.GetAllTableName()
	if err != nil {
		global.RLB_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithDetailed(list, "获取成功", c)
}

/**
 * @Author: Yang
 * @description: 获取表中所有的字段
 * @param {*gin.Context} c
 * @return {*}
 */
func (a *AutoCodeApi) GetFieldsByTableName(c *gin.Context) {
	tableName := c.Query("tableName")
	if tableName == "" {
		response.FailWithMessage("参数错误", c)
		return
	}

	list, err := autoCodeService.GetFieldsByTableName(tableName)
	if err != nil {
		global.RLB_LOG.Error("获取失败!", zap.Error(err))
	}
	response.OkWithDetailed(list, "获取成功", c)
}
