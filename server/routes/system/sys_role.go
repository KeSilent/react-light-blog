/*
 * @Author: Yang
 * @Date: 2025-03-19 21:25:16
 * @Description: 请填写简介
 */
package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/middleware"
)

type RoleRouter struct{}

func (s *RoleRouter) InitRoleRouter(Router *gin.RouterGroup) (R gin.IRoutes) {

	authorityRouter := Router.Group("role").Use(middleware.OperationRecord())
	{
		authorityRouter.GET("getRoleList", roleApi.GetRoleList)
		authorityRouter.GET("getRoleMenus", roleApi.GetRoleMenus)
		authorityRouter.POST("addRoleMenu", roleApi.AddRoleMenu)
	}
	return authorityRouter
}
