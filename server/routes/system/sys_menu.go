package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/middleware"
)

type MenuRouter struct{}

func (s *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	authorityRouter := Router.Group("menu").Use(middleware.OperationRecord())
	{
		authorityRouter.GET("list", menuApi.GetMenuList)
	}
	return authorityRouter
}
