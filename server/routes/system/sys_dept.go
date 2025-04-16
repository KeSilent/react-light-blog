/*
 * @Author: Yang
 * @Date: 2025-04-15 18:24:43
 * @Description: 部门路由
 */
package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/middleware"
)

type DeptRouter struct{}

func (d *DeptRouter) InitDeptRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	deptRouter := Router.Group("dept").Use(middleware.OperationRecord())
	{
		deptRouter.GET("getList", deptApi.GetList)
	}
	return deptRouter
}
