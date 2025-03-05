package system

import (
	"github.com/gin-gonic/gin"
	"github.com/kesilent/react-light-blog/middleware"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("base")
	{
		baseRouter.POST("captcha", baseApi.Captcha)
		baseRouter.POST("login", baseApi.Login)
		baseRouter.POST("initData", baseApi.InitData)
	}
	authorityRouter := Router.Group("authority").Use(middleware.OperationRecord())
	{
		authorityRouter.GET("menus", menuApi.Menus)
	}
	return baseRouter
}
