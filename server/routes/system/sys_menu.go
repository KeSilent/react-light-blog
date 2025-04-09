/*
 * @Author: Yang
 * @Date: 2025-03-19 22:44:20
 * @Description: 请填写简介
 */
package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	authorityRouter := Router.Group("menu").Use(middleware.OperationRecord())
	{
		authorityRouter.GET("getMenuByKey", menuApi.GetMenuByKey)
		authorityRouter.GET("getMenuListByPage", menuApi.GetMenuListByPage)
		authorityRouter.POST("saveBaseMenu", menuApi.SaveBaseMenu)
	}
	return authorityRouter
}
